package main

import (
	"fmt"
	"net/http"
	"time"
)

var directory = http.Dir(".")

func index(w http.ResponseWriter, r *http.Request) {
	file, err := directory.Open("index.html")

	if err != nil {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	defer file.Close()

	w.Header().Set("Content-Type", "text/html")
	http.ServeContent(w, r, "", time.Now(), file)
}

func octet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write([]byte{0xFF, 0x00, 0x00, 0xFF})
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/octet", octet)

	fmt.Println("Listening on 127.0.0.1:8000")

	http.ListenAndServe("127.0.0.1:8000", nil)
}
