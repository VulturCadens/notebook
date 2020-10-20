package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	width  = 800
	height = 600
)

type application struct {
	box *ebiten.Image
}

func (app *application) Update() error {
	return nil
}

func (app *application) Draw(screen *ebiten.Image) {
	c := color.RGBA{50, 50, 50, 255} // https://golang.org/pkg/image/color/
	screen.Fill(c)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(width/4, height/2)
	screen.DrawImage(app.box, options)
}

func (app *application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	app := &application{}

	file, err := os.Open("box.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
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
