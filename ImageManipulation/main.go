package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const (
	width  = 640
	height = 360
)

func main() {
	var i *image.RGBA = image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	var c color.RGBA = color.RGBA{0xCC, 0xCC, 0xCC, 0xFF}

	/*
	 * https://pkg.go.dev/image/draw
	 */

	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.Point{0, 0}, draw.Src)

	file, err := os.Create("image.png")

	defer file.Close()

	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	if err := png.Encode(file, i); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
}
