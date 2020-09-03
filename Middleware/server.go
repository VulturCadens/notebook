package main

import (
	"fmt"
	"net/http"
)

func mainHandler() http.HandlerFunc {
	fmt.Println("One-time (func mainHandler) initialisation.")

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Main Handler")

		w.Write([]byte("<p>(3) Sed eiusmod tempor incidunt ut labore.</p>"))
	}
}

func firstMiddleware(h http.HandlerFunc) http.HandlerFunc {
	fmt.Println("One-time (func firstMiddleware) initialisation.")

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("First Middleware -> ")

		w.Write([]byte("<p>(1) Lorem ipsum dolor sit amet.</p>"))
		h(w, r)
	}
}

func secondMiddleware(h http.HandlerFunc) http.HandlerFunc {
	fmt.Println("One-time (func secondMiddleware) initialisation.")

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Second Middleware -> ")

		w.Write([]byte("<p>(2) Consectetur adipisci elit.</p>"))
		h(w, r)
	}
}

func main() {
	http.HandleFunc("/", firstMiddleware(secondMiddleware(mainHandler())))

	http.ListenAndServe("localhost:8000", nil)
}
