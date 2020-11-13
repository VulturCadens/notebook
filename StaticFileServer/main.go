package main

// cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./www/javascript/

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
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

		// Test if the wasm files have been compressed to brotli format and accepted.
		if path.Ext(p) == ".wasm" {
			acceptEncoding := strings.Split(r.Header.Get("Accept-Encoding"), ",")

			for _, accept := range acceptEncoding {
				if strings.Trim(accept, " ") == "br" {
					file, err := dir.Open(p + ".br")

					if err != nil {
						http.Error(w, "404", http.StatusNotFound)
						return
					}

					defer file.Close()

					w.Header().Set("Content-Type", "application/wasm")
					w.Header().Set("Content-Encoding", "br")
					http.ServeContent(w, r, "", time.Now(), file)

					return
				}
			}

			file, err := dir.Open(p)

			if err != nil {
				http.Error(w, "404", http.StatusNotFound)
				return
			}

			defer file.Close()

			w.Header().Set("Content-Type", "application/wasm")
			http.ServeContent(w, r, "", time.Now(), file)

			return
		}

		// Because a file server would redirect any request ending in "/index.html" anyway...
		if strings.HasSuffix(p, "/index.html") {
			file, err := dir.Open(p)

			if err != nil {
				http.Error(w, "404", http.StatusNotFound)
				return
			}

			defer file.Close()

			/*
			 * If the response's Content-Type header is not set, ServeContent first
			 * tries to deduce the type from name's file extension and, if that fails,
			 * falls back to reading the first block of the content and passing
			 * it to DetectContentType. The name is otherwise unused; in particular
			 * it can be empty and is never sent in the response.
			 *
			 * https://golang.org/pkg/net/http/#ServeContent
			 */
			w.Header().Set("Content-Type", "text/html")
			http.ServeContent(w, r, "", time.Now(), file)

			return
		}

		// - Type http.File
		// https://golang.org/pkg/net/http/#File
		file, err := dir.Open(p)

		if err != nil {
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		defer file.Close()

		// - Type os.FileInfo
		// https://golang.org/pkg/os/#FileInfo
		info, err := file.Stat()

		if err != nil {
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		if info.IsDir() {
			file, err := dir.Open(path.Join(p, "index.html"))

			if err != nil {
				http.Error(w, "404", http.StatusNotFound)
				return
			}

			defer file.Close()

			w.Header().Set("Content-Type", "text/html")
			http.ServeContent(w, r, "", time.Now(), file)

			return
		}

		/*
		 * As a special case, the returned file server redirects any request
		 * ending in "/index.html" to the same path, without the final "index.html".
		 *
		 * https://golang.org/pkg/net/http/#FileServer
		 */
		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", staticFileServer("./www"))

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

	fmt.Printf("Listening on %s\n", server.Addr)

	// https://en.wikipedia.org/wiki/Signal_(IPC)
	// https://golang.org/pkg/os/signal/

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Print("\nShutting down...\n")

	context, cancel := context.WithTimeout(context.Background(), (time.Second * 10))
	defer cancel()

	server.Shutdown(context)
}
