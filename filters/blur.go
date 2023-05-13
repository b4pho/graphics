package filters

import (
	"image"
	"image/color"
	"math"
	"sync"
)

type BlurFilter struct {
	Radius int
}

func (f BlurFilter) Transform(img base) image.Image {
	bounds := img.source.Bounds()
	img2 := image.NewRGBA(bounds)
	kernel := gaussianKernel(bounds, f.Radius)
	var wg sync.WaitGroup
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			wg.Add(1)
			go blurWorker(kernel, img.source, img2, x, y, &wg)
		}
	}
	wg.Wait()
	return img2
}

func averageColor(kernel [][]float64, img image.Image, px0 int, py0 int) color.RGBA {
	imageBounds := img.Bounds()
	kernelSize := len(kernel)
	center := kernelSize / 2
	var avgRed, avgGreen, avgBlue float64
	for y := 0; y < kernelSize; y++ {
		for x := 0; x < kernelSize; x++ {
			px := px0 + (x - center)
			py := py0 + (y - center)
			if px < 0 || py < 0 || px > imageBounds.Max.X || py > imageBounds.Max.Y {
				continue
			}
			red, green, blue, _ := img.At(px, py).RGBA()
			weight := kernel[y][x]
			avgRed += float64(red) * weight
			avgGreen += float64(green) * weight
			avgBlue += float64(blue) * weight
		}
	}
	return color.RGBA{uint8(avgRed / 256), uint8(avgGreen / 256), uint8(avgBlue / 256), uint8(255)}
}

func blurWorker(kernel [][]float64, img image.Image, img2 *image.RGBA, x int, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	avgColor := averageColor(kernel, img, x, y)
	img2.Set(x, y, avgColor)
}

func gaussianFormula(sigma float64, x, y int) float64 {
	return 1 / (2 * math.Pi * sigma * sigma) * math.Exp(-float64(x*x+y*y)/(2*sigma*sigma))
}

func gaussianKernel(bounds image.Rectangle, radius int) [][]float64 {
	size := 2*radius + 1
	sigma := math.Max(float64(radius/2), float64(1))
	sum := float64(0.0)
	matrix := make([][]float64, size)
	for y := -radius; y <= radius; y++ {
		matrix[y+radius] = make([]float64, size)
		for x := -radius; x <= radius; x++ {
			matrix[y+radius][x+radius] = gaussianFormula(sigma, x, y)
			sum += matrix[y+radius][x+radius]
		}

	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			matrix[y][x] /= sum
		}

	}
	return matrix
}
