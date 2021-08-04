package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	width  = 640
	height = 360

	fontFileName = "font/Rajdhani.ttf"
	fontSize     = 30
)

func main() {

	/*
	 *  Create an image.
	 */

	var i *image.RGBA = image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	var c color.RGBA = color.RGBA{0xEE, 0xEE, 0xEE, 0xFF}

	// https://pkg.go.dev/image/draw

	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.Point{0, 0}, draw.Src)

	/*
	 *  Load a font and draw a string.
	 */

	fontData, err := ioutil.ReadFile(fontFileName)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fontParsed, err := truetype.Parse(fontData)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fontDrawer := &font.Drawer{
		Dst: i,
		Src: image.Black,
		Face: truetype.NewFace(fontParsed, &truetype.Options{
			Size:    fontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		}),
	}

	// golang.org/x/image/math/fixed

	fontDrawer.Dot = fixed.Point26_6{
		X: fixed.I(50),
		Y: fixed.I(100),
	}

	fontDrawer.DrawString("Lorem ipsum dolor sit amet.")

	/*
	 *  Save an image.
	 */

	file, err := os.Create("image.png")

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer file.Close()

	if err := png.Encode(file, i); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
