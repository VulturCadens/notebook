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
	positionX float64
	positionY float64
	rotation  float64
	boxImage  *ebiten.Image
	boxSizeX  float64
	boxSizeY  float64
}

func (app *application) Update() error {
	app.rotation += 0.02
	return nil
}

func (app *application) Draw(screen *ebiten.Image) {
	c := color.RGBA{50, 50, 50, 255} // https://golang.org/pkg/image/color/
	screen.Fill(c)

	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(-(app.boxSizeX / 2), -(app.boxSizeY / 2))
	options.GeoM.Rotate(app.rotation)
	options.GeoM.Translate(app.positionX, app.positionY)

	screen.DrawImage(app.boxImage, options)
}

func (app *application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(boxPNG))
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		positionX: 400,
		positionY: 200,
		rotation:  0,
		boxImage:  ebiten.NewImageFromImage(img),
		boxSizeX:  64,
		boxSizeY:  64,
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Window Title")

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
