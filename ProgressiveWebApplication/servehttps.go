package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"
)

const addr = "127.0.0.1:4443"

func staticFileServer(directory string) http.Handler {
	dir := http.Dir(directory)

	fileServer := http.FileServer(dir)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := path.Clean(r.URL.Path)

		if p == "/" || p == "/index.html" {
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

	fmt.Printf("Server is listening on %s \n", addr)

	log.Fatal(http.ListenAndServeTLS(addr, "./WWW/localhost.pem", "./WWW/localhost-key.pem", nil))
}
