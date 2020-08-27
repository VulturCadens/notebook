package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const socket = "/tmp/vultur.sock"

var (
	conn net.Conn
	str  string
	err  error
)

func main() {

	// - Stream      "unix"        SOCK_STREAM (compare to TCP)
	// - Datagram    "unixgram"    SOCK_DGRAM (compare to UDP)
	// - Sequential  "unixpacket"  SOCK_SEQPACKET (compare to SCTP)
	// https://golang.org/src/net/unixsock_posix.go#L16
	//
	// - Type net.Conn interface.
	// https://golang.org/pkg/net/#Conn

	if conn, err = net.Dial("unix", socket); err != nil {
		log.Fatal("Connection error:", err)
	}

	defer conn.Close()

	go func() {
		buffer := make([]byte, 1024)

		for {
			n, err := conn.Read(buffer[:])
			if err != nil {

				if errors.Is(err, io.EOF) {
					fmt.Println("\n !! The server went away. And there's nothing a client can do about it.")
					os.Exit(0)
				} else {
					log.Fatal("Read error:", err)
				}

			}

			fmt.Println("Client got:", string(buffer[0:n]))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a string: ")

		if scanner.Scan() {
			str = scanner.Text()
		}

		if _, err = conn.Write([]byte(str)); err != nil {
			log.Fatal("Write error:", err)
		}

		time.Sleep(250 * time.Millisecond)
	}

}
