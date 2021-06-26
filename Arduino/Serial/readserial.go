package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

const (
	beginByte byte = 0x5B // [
	endByte   byte = 0x5D // ]
)

var (
	value uint16
	bytes []byte       = make([]byte, 4)
	mode  *serial.Mode = &serial.Mode{
		BaudRate: 9600,
	}
)

func main() {
	port, err := serial.Open("/dev/ttyACM0", mode)

	if err != nil {
		log.Fatal(err)
	}

	defer port.Close()

	time.Sleep(time.Second) // Wait to reset the Arduino Micro.

	go func() {
		for {
			if _, err := port.Write([]byte("[ON]")); err != nil {
				log.Fatalln(err)
			}

			time.Sleep(500 * time.Millisecond)

			if _, err := port.Write([]byte("[OFF]")); err != nil {
				log.Fatalln(err)
			}

			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		_, err := port.Read(bytes)

		if err != nil {
			log.Fatalln(err)
		}

		if bytes[0] == beginByte && bytes[3] == endByte {
			value = binary.BigEndian.Uint16(bytes[1:3])

			fmt.Printf("\r [%04d] ", value)
		}
	}
}
