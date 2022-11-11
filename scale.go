package main

import (
	"image"
	"io"
	"math"

	"golang.org/x/image/draw"
)

func ScaleDown(r io.Reader, maxWidth, maxHeight int) (image.Image, string, error) {
	src, format, err := image.Decode(r)
	if err != nil {
		return nil, "", err
	}

	srcWidth := src.Bounds().Dx()
	srcHeight := src.Bounds().Dy()

	// don't resize if source is already smaller
	if srcWidth <= maxWidth || srcHeight <= maxHeight {
		return src, format, err
	}

	// determine scaling ratio
	var ratio float64
	switch {
	case maxWidth == 0:
		ratio = float64(maxHeight) / float64(srcHeight)
	case maxHeight == 0:
		ratio = float64(maxWidth) / float64(srcWidth)
	default:
		ratio = math.Min(
			float64(maxWidth)/float64(srcWidth),
			float64(maxHeight)/float64(srcHeight),
		)
	}

	// scaled down dimension
	newWidth := int(math.Round(float64(srcWidth) * ratio))
	newHeight := int(math.Round(float64(srcHeight) * ratio))

	// Resize:
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	return dst, format, err
}
