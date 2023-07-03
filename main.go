package main

import (
	"flag"
	"fmt"
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

	if inFileName == "" || outFileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := ScaleImage(inFileName, outFileName, width, height)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("scaled %q to %dx%d\n", inFileName, width, height)
}

// ScaleImage resizes the image specified by inFileName and saves it to outFileName with the provided width and height.
func ScaleImage(inFileName, outFileName string, width, height int) error {
	input, err := os.Open(inFileName)
	if err != nil {
		return err
	}
	defer input.Close()

	img, err := imaging.Decode(input)
	if err != nil {
		return err
	}

	// Scale the image if width or height is provided
	if width > 0 || height > 0 {
		img = imaging.Resize(img, width, height, imaging.Lanczos)
	}

	output, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer output.Close()

	err = imaging.Encode(output, img, imaging.PNG)
	if err != nil {
		return err
	}

	return nil
}
