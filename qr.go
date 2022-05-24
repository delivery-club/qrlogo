package qrlogo

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	qr "github.com/skip2/go-qrcode"
)

// Encode encodes QR image, adds logo overlay and renders result as PNG.
func Encode(content string, logo image.Image, size int, outputFormat string) (*bytes.Buffer, error) {
	var buf bytes.Buffer

	code, err := qr.New(content, qr.Highest)
	if err != nil {
		return nil, err
	}

	img := code.Image(size)
	p, ok := img.(*image.Paletted)
	if !ok {
		return nil, fmt.Errorf("undefined qr-code type provided: %T", img)
	}

	switch outputFormat {
	case "png":
		if err = png.Encode(&buf, overlayLogo(p, logo)); err != nil {
			return nil, err
		}
	case "jpeg", "jpg":
		if err = jpeg.Encode(&buf, overlayLogo(p, logo), nil); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", outputFormat)
	}

	return &buf, nil
}

// overlayLogo - blends logo to the center of the QR code.
func overlayLogo(dst *image.Paletted, logo image.Image) *image.NRGBA {
	res := image.NewNRGBA(dst.Rect)

	for x := 0; x < dst.Bounds().Max.X; x++ {
		for y := 0; y < dst.Bounds().Max.Y; y++ {
			res.Set(x, y, dst.At(x, y))
		}
	}

	offsetX := dst.Bounds().Max.X/2 - logo.Bounds().Max.X/2
	offsetY := dst.Bounds().Max.Y/2 - logo.Bounds().Max.Y/2
	for x := 0; x < logo.Bounds().Max.X; x++ {
		for y := 0; y < logo.Bounds().Max.Y; y++ {
			res.Set(x+offsetX, y+offsetY, logo.At(x, y))
		}
	}

	return res
}
