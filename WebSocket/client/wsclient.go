package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Connecting...")

	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8000/", nil)

	if err != nil {
		log.Fatal("Connection error:", err)
	}

	fmt.Println("Connection OK")

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Print("Read error:", err)
			break
		}

		fmt.Println("Message received:", string(message))
	}
}
