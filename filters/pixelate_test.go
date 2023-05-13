package filters

import (
    "testing"
)

func TestPixelateFilter(t *testing.T) {
    input, err := OpenImage("../assets/original.jpg")
    if err != nil {
        t.Errorf("Error, cannot load original image: %s", err)
    }
    
    output, err := OpenImage("../assets/pixelate_10.png")
    if err != nil {
       t.Errorf("Error, cannot load output image: %s", err)
    }
    
    f := PixelateFilter{10}
    img2 := f.Transform(input)
    if !assertSamePicture(output.source, img2) {
        t.Error("Regression failed: new image generated is not the same of output image")
    }
}

func TestPixelateFilterWithDifferentParam(t *testing.T) {
    input, err := OpenImage("../assets/original.jpg")
    if err != nil {
        t.Errorf("Error, cannot load original image: %s", err)
    }
    
    output, err := OpenImage("../assets/pixelate_20.png")
    if err != nil {
       t.Errorf("Error, cannot load output image: %s", err)
    }
    
    f := PixelateFilter{20}
    img2 := f.Transform(input)
    if !assertSamePicture(output.source, img2) {
        t.Error("Regression failed: new image generated is not the same of output image")
    }
}
