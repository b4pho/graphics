package main

import (
	"flag"
	"graphics/filters"
	"log"
)

func main() {
	inputFilename := flag.String("source", "input.png", "source filename, it could be any image file (jpg, png, bmp, gif)")
	outputFilename := flag.String("destination", "output.png", "destination filename, it could be any image file (jpg, png, bmp, gif)")
	filter := flag.String("filter", "blur", "choose among: grayscale, blur, sobel, inverted, pixelate")
	radius := flag.Int("radius", 5, "this option is related to blur filter")
	pixelSize := flag.Int("pixel-size", 5, "this option is related to pixelate filter")
	flag.Parse()

	switch *filter {
	case "grayscale":
		doFilter(*inputFilename, *outputFilename, filters.GrayscaleFilter{})
	case "blur":
		doFilter(*inputFilename, *outputFilename, filters.BlurFilter{Radius: *radius})
	case "sobel":
		doFilter(*inputFilename, *outputFilename, filters.SobelFilter{})
	case "inverted":
		doFilter(*inputFilename, *outputFilename, filters.InvertFilter{})
	case "pixelate":
		doFilter(*inputFilename, *outputFilename, filters.PixelateFilter{PixelSize: *pixelSize})
	default:
		log.Fatal("I dont know this filter: ", *filter)
	}
}

func doFilter(source, destination string, f filters.Filter) {
	img, err := filters.OpenImage(source)
	if err != nil {
		log.Fatal(err)
	}
	err = img.ApplyFilter(f)
	if err != nil {
		log.Fatal(err)
	}
	err = img.StoreImage(destination)
	if err != nil {
		log.Fatal(err)
	}
}
