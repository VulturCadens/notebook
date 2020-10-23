package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	wait = 500
	pin  = rpio.Pin(26)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Output()

	for {
		pin.Toggle()

		time.Sleep(time.Millisecond * wait)
	}
}
