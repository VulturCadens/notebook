package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed box.png
var boxPNG []byte

const (
	width  = 1280
	height = 720
)

type sprite struct {
	positionX float64
	positionY float64
	rotation  float64
	speed     float64
	image     *ebiten.Image
}

type application struct {
	sprites []sprite
}

func (app *application) Update() error {
	for i := range app.sprites {
		app.sprites[i].rotation += app.sprites[i].speed
	}

	return nil
}

func (app *application) Draw(screen *ebiten.Image) {
	c := color.RGBA{50, 50, 50, 255} // https://golang.org/pkg/image/color/
	screen.Fill(c)

	for i := range app.sprites {
		options := &ebiten.DrawImageOptions{}

		options.GeoM.Translate(-32, -32)
		options.GeoM.Rotate(app.sprites[i].rotation)
		options.GeoM.Translate(app.sprites[i].positionX, app.sprites[i].positionY)

		screen.DrawImage(app.sprites[i].image, options)
	}
}

func (app *application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(boxPNG))
	if err != nil {
		log.Fatal(err)
	}

	app := &application{}

	for i := 0; i < 50; i++ {
		s := sprite{
			positionX: float64(rand.Intn(1200-10) + 10),
			positionY: float64(rand.Intn(700-10) + 10),
			rotation:  float64(rand.Intn(3 - 0)),
			speed:     rand.Float64() / 10,
			image:     ebiten.NewImageFromImage(img),
		}

		app.sprites = append(app.sprites, s)
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Window Title")

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
