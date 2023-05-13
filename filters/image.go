package filters

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type Filter interface {
	Transform(img base) image.Image
}

type base struct {
	source      image.Image
	destination image.Image
}

func OpenImage(filename string) (base, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return base{}, err
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return base{}, err
	}
	return base{img, nil}, nil
}

func (b *base) ApplyFilter(f Filter) error {
	if b.source == nil {
		return errors.New("no image to transform")
	}
	b.destination = f.Transform(*b)
	return nil
}

func (b base) StoreImage(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if b.destination == nil {
		return errors.New("no transformed image to store")
	}
	tokens := strings.Split(strings.ToLower(filename), ".")
	fileExtension := tokens[len(tokens)-1]
	switch fileExtension {
	case "png":
		png.Encode(f, b.destination)
	case "gif":
		gif.Encode(f, b.destination, nil)
	case "jpg":
		jpeg.Encode(f, b.destination, nil)
	case "jpeg":
		jpeg.Encode(f, b.destination, nil)
	default:
		return errors.New("cannot save file - uknown extension: " + fileExtension)
	}
	return nil
}
