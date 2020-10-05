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

	w.Header().Set("Content-Type", "text/html")
	http.ServeContent(w, r, "", time.Now(), file)
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Listening on 127.0.0.1:8000...")

	http.ListenAndServe("127.0.0.1:8000", nil)
}
