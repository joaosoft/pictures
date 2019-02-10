package pictures

import (
	"image"
	"image/draw"
	"os"
)

type Pictures struct {
	image.Image
}

func Open(filename string) (*Pictures, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return &Pictures{Image: img}, nil
}

func (p *Pictures) Save(filename string, encoder Encoder) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return encoder(f, p.Image)
}

func (p *Pictures) Crop(x0, y0, x1, y1 int) *Pictures {
	bounds := p.Image.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, p.Image, bounds.Min, draw.Src)

	p.Image = imgRGBA.SubImage(image.Rect(x0, y0, x1, y1))

	return p
}

func (p *Pictures) Resize(width, height int) *Pictures {
	if width <= 0 || height <= 0 || p.Image.Bounds().Empty() {
		p.Image = image.NewRGBA(image.Rect(0, 0, 0, 0))
		return p
	}

	bounds := p.Image.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, p.Image, bounds.Min, draw.Src)

	imgRGBAW, imgRGBAH := imgRGBA.Bounds().Dx(), imgRGBA.Bounds().Dy()
	imgRGBAStride := imgRGBA.Stride

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dstStride := dst.Stride

	dx := float64(imgRGBAW) / float64(width)
	dy := float64(imgRGBAH) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := y*dstStride + x*4
			ipos := int((float64(y)+0.5)*dy)*imgRGBAStride + int((float64(x)+0.5)*dx)*4

			dst.Pix[pos+0] = imgRGBA.Pix[ipos+0]
			dst.Pix[pos+1] = imgRGBA.Pix[ipos+1]
			dst.Pix[pos+2] = imgRGBA.Pix[ipos+2]
			dst.Pix[pos+3] = imgRGBA.Pix[ipos+3]
		}
	}

	p.Image = dst

	return p
}