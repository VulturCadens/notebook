package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	host = "127.0.0.1"
	port = "8000"
)

type state struct {
	Pin   int `json:"pin"`
	State int `json:"state"`
}

//go:embed content/index.html
var indexhtml []byte

func command(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	state := state{}

	err = json.Unmarshal(payload, &state)

	if err != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	fmt.Println(state)

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"OK\"}"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write(indexhtml)
}

func main() {
	http.HandleFunc("/command", command)
	http.HandleFunc("/", index)

	addr := net.JoinHostPort(host, port)

	server := &http.Server{
		Addr:           addr,
		Handler:        nil,
		WriteTimeout:   time.Second * 5,
		ReadTimeout:    time.Second * 5,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 32768,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {

			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}

		}
	}()

	fmt.Printf("Listening on %s\n", server.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Print("\nShutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 10))
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	fmt.Print("ok\n\n")
}
