package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

var (
	count     int
	templates = template.Must(template.ParseFiles("./www/index.html"))
)

func init() {
	file, err := os.Open("count.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := scanner.Text()

	file.Close()

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.Create("count.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i++
	s = strconv.Itoa(i)

	if _, err := file.WriteString(s); err != nil {
		log.Fatal(err)
	}

	count = i
	log.Printf("[%d]\n", i)
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
	content := &struct{ Count int }{count}

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
