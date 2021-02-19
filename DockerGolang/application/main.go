package main

import (
	"net/http"
	"os"
	"time"
)

func favicon(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./www/favicon-32x32.png")

	if err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	w.Header().Set("Content-Type", "image/png")
	http.ServeContent(w, r, "", time.Now(), file)
}

func response(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./www/index.html")

	if err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeContent(w, r, "", time.Now(), file)
}

func main() {
	http.HandleFunc("/favicon-32x32.png", favicon)
	http.HandleFunc("/", response)

	http.ListenAndServe(":8000", nil)
}
