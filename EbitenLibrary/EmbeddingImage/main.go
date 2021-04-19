package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed box.png
var boxPNG []byte

const (
	width  = 1280
	height = 720
)

type application struct {
	x   float64
	y   float64
	box *ebiten.Image
}

func (app *application) Update() error {
	return nil
}

func (app *application) Draw(screen *ebiten.Image) {
	c := color.RGBA{50, 50, 50, 255} // https://golang.org/pkg/image/color/
	screen.Fill(c)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(app.x, app.y)
	screen.DrawImage(app.box, options)
}

func (app *application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	app := &application{
		x: width / 2,
		y: height / 2,
	}

	img, _, err := image.Decode(bytes.NewReader(boxPNG))
	if err != nil {
		log.Fatal(err)
	}

	app.box = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Window Title")

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
