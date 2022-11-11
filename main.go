package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
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

	img, format, err := ScaleDown(input, width, height)
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

	switch format {
	case "jpeg":
		jpeg.Encode(output, img, nil)
	case "png":
		png.Encode(output, img)
	}

	fmt.Printf("scaled %s %q to %dx%d\n",
		format, inFileName, img.Bounds().Dx(), img.Bounds().Dy())
}
