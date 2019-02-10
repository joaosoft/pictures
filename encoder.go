package pictures

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/bmp"
)

type Encoder func(io.Writer, image.Image) error

func JPEGEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}

func PNGEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}

func BMPEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return bmp.Encode(w, img)
	}
}
