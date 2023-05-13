package filters

import (
    "testing"
    "image"
)

func TestGrayscaleFilter(t *testing.T) {
    input, err := OpenImage("../assets/original.jpg")
    if err != nil {
        t.Errorf("Error, cannot load original image: %s", err)
    }
    
    output, err := OpenImage("../assets/grayscale.png")
    if err != nil {
       t.Errorf("Error, cannot load output image: %s", err)
    }
    
    f := GrayscaleFilter{}
    img2 := f.Transform(input)
    if !assertSamePicture(output.source, img2) {
        t.Error("Regression failed: new image generated is not the same of output image")
    }
}

func assertSamePicture(a, b image.Image) bool {
    if a.Bounds().Max.X != b.Bounds().Max.X || a.Bounds().Max.Y != b.Bounds().Max.Y {
        return false
    }
    for y := 0; y < a.Bounds().Max.Y; y++ {
        for x := 0; x < a.Bounds().Max.X; x++ {
            r1, g1, b1, a1 := a.At(x, y).RGBA()
            r2, g2, b2, a2 := b.At(x, y).RGBA()
            if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
                return false
            }
        }
    }
    return true
}
