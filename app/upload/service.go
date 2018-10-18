package upload

import (
	"image/jpeg"
	"image/png"
	"io"

	"github.com/chai2010/webp"
)

// SaveAvator ...
func SaveAvator() {

}

// Png2Webp ...
func Png2Webp(source io.Reader, target io.Writer) error {
	img, err := png.Decode(source)
	if err != nil {
		return err
	}
	err = webp.Encode(target, img, &webp.Options{Lossless: true})
	return err
}

// Jpg2Webp ...
func Jpg2Webp(source io.Reader, target io.Writer) error {
	img, err := jpeg.Decode(source)
	if err != nil {
		return err
	}
	err = webp.Encode(target, img, &webp.Options{Lossless: true})
	return err
}
