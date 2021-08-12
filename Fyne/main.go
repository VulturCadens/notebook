package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type countButton struct {
	counter int
	button  *widget.Button
	label   *widget.Label
}

func (b *countButton) click() {
	b.counter++
	b.label.SetText(fmt.Sprintf("Button is clicked %d times.", b.counter))
}

func newCountButton(title string, label *widget.Label) countButton {
	var b countButton = countButton{
		counter: 0,
		label:   label,
	}

	b.button = widget.NewButton(title, func() { b.click() })

	return b
}

func main() {
	application := app.New()

	window := application.NewWindow("Hello World")

	textLabel := widget.NewLabel("This is label.")
	center := container.New(layout.NewCenterLayout(), textLabel)

	s := binding.NewString()
	tickLabel := widget.NewLabelWithData(s)

	spacer := layout.NewSpacer()
	clickButton := newCountButton("Click Me", textLabel)

	window.SetContent(container.NewVBox(
		tickLabel,
		center,
		spacer,
		clickButton.button,
	))

	window.Resize(fyne.NewSize(600, 300))
	window.Show()

	i := 0

	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			i++
			s.Set(fmt.Sprintf("Ticks: %d", i))
		}
	}()

	application.Run()
}
