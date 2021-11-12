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

func loadImage(fileName string) image.Image {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer file.Close()

	image, _, err := image.Decode(file)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return image
}

func main() {

	/*
	 *  Create an image.
	 */

	const (
		width  int = 640
		height int = 360
	)

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

	const (
		fontFileName string  = "font/Rajdhani.ttf"
		fontSize     float64 = 30
	)

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

	// https://pkg.go.dev/golang.org/x/image/math/fixed

	fontDrawer.Dot = fixed.Point26_6{
		X: fixed.I(50),
		Y: fixed.I(100),
	}

	fontDrawer.DrawString("Lorem ipsum dolor sit amet.")

	/*
	 *  Read an image from the file and combine images.
	 */

	ship := loadImage("ship.png")

	const (
		targetX int = 100
		targetY int = 250
		size    int = 64
	)

	draw.Draw(
		i,
		image.Rectangle{
			image.Point{targetX, targetY},
			image.Point{targetX + size, targetY + size},
		},
		ship,
		image.Point{0, 0},
		draw.Over,
	)

	/*
	 *  Save the image.
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
