package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	var inFileName, outFileName string
	var width, height int

	flag.StringVar(&inFileName, "in", "", "input image filename")
	flag.StringVar(&outFileName, "out", "", "output image filename")
	flag.IntVar(&width, "width", 0, "new width or 0")
	flag.IntVar(&height, "height", 0, "new height or 0")

	flag.Parse()

	if (inFileName == "") || (outFileName == "") {
		flag.Usage()
		os.Exit(1)
	}

	input, err := os.Open(inFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer input.Close()

	img, err := ScaleDown(input, width, height)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output, err := os.Create(outFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer output.Close()

	format, err := imaging.FormatFromFilename(outFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch format {
	case imaging.JPEG:
		imaging.Encode(output, img, imaging.JPEG, imaging.JPEGQuality(95))

	case imaging.PNG:
		imaging.Encode(output, img, imaging.PNG, imaging.PNGCompressionLevel(png.DefaultCompression))
	}

	fmt.Printf("scaled %q to %dx%d\n",
		inFileName, img.Bounds().Dx(), img.Bounds().Dy())
}
