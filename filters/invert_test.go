package filters

import (
    "testing"
)

func TestInvertFilter(t *testing.T) {
    input, err := OpenImage("../assets/original.jpg")
    if err != nil {
        t.Errorf("Error, cannot load original image: %s", err)
    }
    
    output, err := OpenImage("../assets/inverted.png")
    if err != nil {
       t.Errorf("Error, cannot load output image: %s", err)
    }
    
    f := InvertFilter{}
    img2 := f.Transform(input)
    if !assertSamePicture(output.source, img2) {
        t.Error("Regression failed: new image generated is not the same of output image")
    }
}
