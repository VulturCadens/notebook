package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
)

//go:embed www
var filesystem embed.FS

const (
	host = "127.0.0.1"
	port = "8000"
)

func staticFileServer(content fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(content))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := path.Clean(r.URL.Path)
		p := strings.TrimPrefix(s, "/")

		if p == "index.html" || len(p) == 0 {
			c, err := fs.ReadFile(content, "index.html")

			if err != nil {
				log.Fatal(err)
			}

			w.Header().Set("Content-Type", "text/html")
			w.Write(c)

			return
		}

		file, err := content.Open(p)

		if err != nil {
			fmt.Println("ERROR: can't open " + p)
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		defer file.Close()

		info, _ := file.Stat()

		if info.IsDir() {
			fmt.Println("ERROR: Is directory " + p)
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	content, err := fs.Sub(filesystem, "www")

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", staticFileServer(content))

	server := &http.Server{
		Addr:           net.JoinHostPort(host, port),
		Handler:        nil,
		MaxHeaderBytes: 32768,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {

			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}

		}
	}()

	fmt.Printf("Running a server on %s\n", server.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Print("\nBye... \n")
}
