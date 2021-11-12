package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Const ...
const (
	DEVICE string = "/dev/input/js0"

	BUTTON uint8 = 0x01
	AXIS   uint8 = 0x02
	INIT   uint8 = 0x80
)

type joystick struct {
	time   uint32
	value  uint16
	event  uint8
	number uint8
}

var (
	file       *os.File
	err        error
	buffer     []byte = make([]byte, 8)
	js         joystick
	valueAsInt int
)

func main() {
	if file, err = os.Open(DEVICE); err != nil {
		panic(err)
	}

	defer file.Close()

	for {
		if _, err = io.ReadFull(file, buffer); err != nil {
			fmt.Println(err)
			break
		}

		js = joystick{
			time:   binary.LittleEndian.Uint32(buffer[0:4]),
			value:  binary.LittleEndian.Uint16(buffer[4:6]),
			event:  buffer[6:7][0],
			number: buffer[7:8][0],
		}

		if js.event == BUTTON {

			fmt.Printf("Time: %d -> ", js.time)

			if js.value == 0 {
				fmt.Printf("Button %d is released \n", js.number)
			} else {
				fmt.Printf("Button %d is pressed \n", js.number)
			}

		} else if js.event == AXIS {

			fmt.Printf("Time: %d -> ", js.time)

			if js.value > (1 << 15) { // 32768
				valueAsInt = int(js.value) - (1 << 16) // 65536
			} else {
				valueAsInt = int(js.value)
			}

			fmt.Printf("Axis %d is %d \n", js.number, valueAsInt)

		} else if js.event == INIT {

			fmt.Printf("Time: %d -> Init Event \n", js.time)

		}
	}

	fmt.Println("Exit...")
}
