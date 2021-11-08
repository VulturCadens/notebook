package main

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"time"
)

var (
	directory = http.Dir(".")
	words     = [2]uint16{255, 65280}
	twoBytes  = make([]byte, 2, 2)
)

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

	for _, word := range words {
		binary.LittleEndian.PutUint16(twoBytes[0:], word)
		w.Write(twoBytes)

		fmt.Printf("Sent %d = %d (little-endian)\n", twoBytes, word)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/octet", octet)

	fmt.Println("Listening on 127.0.0.1:8000")

	http.ListenAndServe("127.0.0.1:8000", nil)
}
