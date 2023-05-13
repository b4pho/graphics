package filters

import (
    "image"
    "image/color"
)

type PixelateFilter struct {
    PixelSize int
}

func (f PixelateFilter) Transform(img base) image.Image  {
    bounds := img.source.Bounds()
    img2 := image.NewRGBA(bounds)
    for y := 0; y < bounds.Max.Y; y++ {
        for x := 0; x < bounds.Max.X; x++ {
            img2.Set(x, y, pixelColor(img, x, y, f.PixelSize))
        }
    } 
    return img2   
}

func pixelColor(img base, x, y, pixelSize int) color.RGBA {
    if pixelSize < 1 {
        pixelSize = 1
    }
    px := x / pixelSize * pixelSize
    py := y / pixelSize * pixelSize
    red, green, blue, alpha := img.source.At(px, py).RGBA()
    return color.RGBA{uint8(red >> 8), uint8(green >> 8), uint8(blue >> 8), uint8(alpha >> 8)}
}
