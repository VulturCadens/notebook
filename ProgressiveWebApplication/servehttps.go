package main

import (
	"log"
	"net"
	"net/http"
	"path"
	"time"
)

const (
	host = "127.0.0.1"
	port = "4443"
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

			defer file.Close()

			w.Header().Set("Content-Type", "text/html")
			http.ServeContent(w, r, "", time.Now(), file)

			return
		}

		file, err := dir.Open(p)

		if err != nil {
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		defer file.Close()

		info, err := file.Stat()

		if err != nil || info.IsDir() {
			http.Error(w, "404", http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", staticFileServer("./WWW"))

	addr := net.JoinHostPort(host, port)

	log.Fatal(http.ListenAndServeTLS(addr, "./WWW/localhost.pem", "./WWW/localhost-key.pem", nil))
}
