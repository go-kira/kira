package upload

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"

	// this for convert any image type to specefic one.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

// convertToJPEG converts from any recognized format to JPEG.
func convertToJPEG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, nil)
}
