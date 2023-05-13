package filters

import (
    "image"
    "image/color"
)

type GrayscaleFilter struct {}

func (f GrayscaleFilter) Transform(img base) image.Image {
    bounds := img.source.Bounds()
    img2 := image.NewRGBA(bounds)
    for y := 0; y < bounds.Max.Y; y++ {
        for x := 0; x < bounds.Max.X; x++ {
            img2.Set(x, y, grayColor(img, x, y))
        }
    } 
    return img2 
}

func grayColor(img base, x, y int) color.RGBA {
    red, green, blue, alpha := img.source.At(x, y).RGBA()
    gray := uint32(0.299 * float64(red) + 0.587 * float64(green) + 0.114 * float64(blue))
    return color.RGBA{uint8(gray >> 8), uint8(gray >> 8), uint8(gray >> 8), uint8(alpha >> 8)}
}
