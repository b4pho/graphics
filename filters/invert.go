package filters

import (
    "image"
    "image/color"
)

type InvertFilter struct {}


func (f InvertFilter) Transform(img base) image.Image  {
    bounds := img.source.Bounds()
    img2 := image.NewRGBA(bounds)
    for y := 0; y < bounds.Max.Y; y++ {
        for x := 0; x < bounds.Max.X; x++ {
            img2.Set(x, y, invertedColor(img, x, y))
        }
    } 
    return img2  
}

func invertedColor(img base, x, y int) color.RGBA {
    red, green, blue, alpha := img.source.At(x, y).RGBA()
    red, green, blue, alpha = 255 - (red >> 8), 255 - (green >> 8), 255 - (blue >> 8), alpha >> 8
    return color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
}
