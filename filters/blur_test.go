package filters

import (
	"testing"
)

func TestBlurFilter(t *testing.T) {
	input, err := OpenImage("../assets/original.jpg")
	if err != nil {
		t.Errorf("error, cannot load original image: %s", err)
	}

	output, err := OpenImage("../assets/blur_10.png")
	if err != nil {
		t.Errorf("error, cannot load output image: %s", err)
	}

	f := BlurFilter{10}
	img2 := f.Transform(input)
	if !assertSamePicture(output.source, img2) {
		t.Error("regression failed: new image generated is not the same of output image")
	}
}

func TestBlurFilterWithDifferentParam(t *testing.T) {
	input, err := OpenImage("../assets/original.jpg")
	if err != nil {
		t.Errorf("error, cannot load original image: %s", err)
	}

	output, err := OpenImage("../assets/blur_5.png")
	if err != nil {
		t.Errorf("error, cannot load output image: %s", err)
	}

	f := BlurFilter{5}
	img2 := f.Transform(input)
	if !assertSamePicture(output.source, img2) {
		t.Error("regression failed: new image generated is not the same of output image")
	}
}
