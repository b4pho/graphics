package filters

import (
    "math"
    "image"    
    "image/color"
)

type SobelFilter struct {}

func (f SobelFilter) Transform(img base) image.Image  {
    bounds := img.source.Bounds()
    img2 := image.NewRGBA(bounds)
    hKernel := sobelKernel(true)
    vKernel := sobelKernel(false)
    for y := 0; y < bounds.Max.Y; y++ {
        for x := 0; x < bounds.Max.X; x++ {
            img2.Set(x, y, edgeColor(hKernel, vKernel, img, x, y))
        }
    } 
    return img2
}

func edgeColor(hKernel, vKernel [][]float64, img base, px0, py0 int) color.RGBA {
    imageBounds := img.source.Bounds()
    kernelSize := 3
    center := kernelSize / 2;
    var hGray, vGray float64;
    for y := 0; y < kernelSize; y++ {
        for x := 0; x < kernelSize; x++ {
            px := px0 + (x - center)
            py := py0 + (y - center)
            gray := float64(0)
            var red, green, blue uint32
            if !(px < 0 || py < 0 || px > imageBounds.Max.X || py > imageBounds.Max.Y) {
                red, green, blue, _ = img.source.At(px, py).RGBA()
                gray = 0.299 * float64(red) + 0.587 * float64(green) + 0.114 * float64(blue) 
                gray = gray / 8
            }
            hWeight := hKernel[y][x] 
            vWeight := vKernel[y][x]
            hGray += gray * hWeight
            vGray += gray * vWeight          
        }
    }
    gray := uint32(math.Sqrt(hGray * hGray + vGray * vGray))
    return color.RGBA{uint8(gray >> 8), uint8(gray >> 8), uint8(gray >> 8), uint8(255)}     
}

func sobelKernel(horizontal bool) [][]float64 {
    size := 3
    matrix := make([][]float64, size)
    for y := 0; y < size; y++ {
        matrix[y] = make([]float64, size)
        for x := 0; x < size; x++ {
            value := float64(0)
            if horizontal {
                if y == 0 {
                    value = float64(-(x + 1))
                } else if y == 2 {
                    value = float64(x + 1)
                }
                matrix[y][x] = value
            } else {
                if x == 0 {
                    value = float64(-(y + 1))
                } else if x == 2 {
                    value = float64(y + 1)
                }                
                matrix[y][x] = value
            }
        }
    }
    return matrix
}
