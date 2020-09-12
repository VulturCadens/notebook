package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"
)

const (
	host = "127.0.0.1"
	port = "8000"
)

func staticFileServer(directory string) http.Handler {
	dir := http.Dir(directory)

	fileServer := http.FileServer(dir)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := path.Clean(r.URL.Path)

		if p == "/" {
			file, err := dir.Open("index.html")

			if err != nil {
				http.Error(w, "404", http.StatusNotFound)
				return
			}

			http.ServeContent(w, r, (directory + "/index.html"), time.Now(), file)

			return
		}

		// - Type http.File
		// https://golang.org/pkg/net/http/#File
		file, err := dir.Open(p)

		if err != nil {
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		// - Type os.FileInfo
		// https://golang.org/pkg/os/#FileInfo
		info, err := file.Stat()

		if err != nil || info.Mode().IsDir() {
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", staticFileServer("./www"))
	http.Handle("/css/", http.StripPrefix("/css/", staticFileServer("./www/css")))

	// fmt.Sprintf("%s:%d", host, port) VS net.JoinHostPort(host, port) !!
	addr := net.JoinHostPort(host, port)

	server := &http.Server{
		Addr:           addr,
		Handler:        nil, // The DefaultServeMux is used.
		WriteTimeout:   time.Second * 10,
		ReadTimeout:    time.Second * 10,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 32768, // 32 KiB
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	fmt.Printf("Listening at %s", server.Addr)

	// https://en.wikipedia.org/wiki/Signal_(IPC)
	// https://golang.org/pkg/os/signal/

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Print("Shutting down...")

	context, cancel := context.WithTimeout(context.Background(), (time.Second * 10))
	defer cancel()

	server.Shutdown(context)

	fmt.Println("Goodbye!")
}
