// go run .

package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	width  = 800
	height = 600
)

var backgroundColor = color.RGBA{50, 50, 50, 255} // https://golang.org/pkg/image/color/

type application struct {
	x          float64
	y          float64
	rot        float64
	speed      float64
	ship       *ebiten.Image
	gamepadIDs map[ebiten.GamepadID]struct{}
}

func (app *application) Update() error {
	if app.gamepadIDs == nil {
		app.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}

	for _, id := range inpututil.JustConnectedGamepadIDs() {
		log.Printf("Connected ID: %d \n", id)
		app.gamepadIDs[id] = struct{}{}
	}

	for id := range app.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("Disconnected ID: %d \n", id)
			delete(app.gamepadIDs, id)
		}
	}

	for id := range app.gamepadIDs {
		valueHorizontal := ebiten.GamepadAxis(id, 3)
		valueVertical := ebiten.GamepadAxis(id, 4)

		if valueHorizontal < -0.1 || valueHorizontal > 0.1 {
			app.rot += valueHorizontal * 0.1
		}

		if valueVertical < -0.1 || valueVertical > 0.1 {
			app.speed += valueVertical * 0.1
		}
	}

	app.x += app.speed * math.Sin(-app.rot)
	app.y += app.speed * math.Cos(-app.rot)

	return nil
}

func (app *application) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-32, -32) // Rotation center of rectangle.
	options.GeoM.Rotate(app.rot)
	options.GeoM.Translate(app.x, app.y)

	screen.DrawImage(app.ship, options)
}

func (app *application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	app := &application{
		x: width / 2,
		y: height / 2,
	}

	file, err := os.Open("ship.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	app.ship = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Window Title")

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
