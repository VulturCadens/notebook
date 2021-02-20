package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var (
	c         uint64
	templates = template.Must(template.ParseFiles("./www/index.html"))
)

func init() {
	var (
		file *os.File
		err  error
	)

	if file, err = os.OpenFile("count.gob", os.O_SYNC|os.O_RDWR, 0644); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file) // read
	encoder := gob.NewEncoder(file) // write

	decoder.Decode(&c)

	c++

	file.Truncate(0)
	file.Seek(0, 0)

	encoder.Encode(c)

	log.Printf("[%d]\n", c)
}

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
	content := &struct{ Count uint64 }{c}

	if err := templates.ExecuteTemplate(w, "INDEX", content); err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/favicon-32x32.png", favicon)
	http.HandleFunc("/", response)

	http.ListenAndServe(":8000", nil)
}
