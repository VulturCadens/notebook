package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const socket = "/tmp/vultur.sock"

var (
	listener   net.Listener
	err        error
	identifier int = 0
)

func main() {

	if err = os.RemoveAll(socket); err != nil {
		log.Fatal(err)
	}

	// - Stream      "unix"        SOCK_STREAM (compare to TCP)
	// - Datagram    "unixgram"    SOCK_DGRAM (compare to UDP)
	// - Sequential  "unixpacket"  SOCK_SEQPACKET (compare to SCTP)
	// https://golang.org/src/net/unixsock_posix.go#L16
	//
	// - Type net.Listener interface.
	// https://golang.org/pkg/net/#Listener

	if listener, err = net.Listen("unix", socket); err != nil {
		log.Fatal("Listen error:", err)
	}

	defer listener.Close()

	for {

		// - Type net.Conn interface.
		// https://golang.org/pkg/net/#Conn

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		identifier++

		go func() {
			defer conn.Close()

			connID := identifier

			for {
				buffer := make([]byte, 1024)

				n, err := conn.Read(buffer)
				if err != nil {

					if errors.Is(err, io.EOF) {
						fmt.Printf(" !! Goodbye, a client (%d) went away... \n", connID)
						break
					} else {
						log.Fatal("Read error:", err)
					}

				}

				fmt.Printf("Client (%d) sent: %s \n", connID, string(buffer[0:n]))

				if _, err = conn.Write(buffer[0:n]); err != nil {
					log.Fatal("Write error:", err)
				}
			}
		}()

	}

}
