package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	randomGenerator     = rand.Reader
	identifier      int = 0
	mutex           sync.Mutex
)

func webSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Connection error:", err)
		return
	}

	defer conn.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	mutex.Lock()
	identifier++
	mutex.Unlock()

	var (
		randomSlice  []byte        = make([]byte, 8)
		randomBase64 []byte        = make([]byte, 12)
		connID       int           = identifier
		quit         chan struct{} = make(chan struct{}) // Allocation isn't necessary (struct{} contains nothing).
	)

	go func() {
		defer close(quit)

		for {
			_, message, err := conn.ReadMessage()

			if err != nil {
				log.Print("Read error:", err)
				return
			}

			fmt.Printf("Message received from the client (%d): %s \n", connID, string(message))
		}
	}()

	for {
		select {

		case <-quit:
			return

		case <-ticker.C:
			if _, err := io.ReadFull(randomGenerator, randomSlice); err != nil {
				log.Print("Random generator error:", err)
				return
			}

			// Package encoding/base64
			// https://golang.org/pkg/encoding/base64/
			base64.StdEncoding.Encode(randomBase64, randomSlice)

			if err := conn.WriteMessage(websocket.TextMessage, randomBase64); err != nil {
				log.Print("Write error:", err)
				return
			}

			fmt.Printf("Sent to the client (%d): %s \n", connID, string(randomBase64))
		}
	}
}

func main() {
	http.HandleFunc("/", webSocketConnection)

	http.ListenAndServe("127.0.0.1:8000", nil)
}
