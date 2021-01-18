package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

const cookieName = "session_example"
const cookieValue = "foobar"

var users = map[string]string{
	"john": "smith",
}

func authentication(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(cookieName)

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(w, "401", http.StatusUnauthorized)
				return
			}

			http.Error(w, "400", http.StatusBadRequest)
			return
		}

		token := c.Value

		if token != cookieValue {
			http.Error(w, "401", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}

func void(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404", http.StatusNotFound)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	if _, ok := users[username]; ok {
		if users[username] == password {
			file, err := os.Open("welcome.html")

			if err != nil {
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}

			defer file.Close()

			http.SetCookie(w, &http.Cookie{
				Name:     cookieName,
				Path:     "/",
				Value:    cookieValue,
				HttpOnly: true,
				// Secure:   true,
				SameSite: http.SameSiteStrictMode,
				Expires:  time.Now().Add(120 * time.Second),
			})

			w.Header().Set("Content-Type", "text/html")
			http.ServeContent(w, r, "", time.Now(), file)
			return
		}
	}

	http.Error(w, "403", http.StatusForbidden)
}

func login(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("login.html")

	if err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	w.Header().Set("Content-Type", "text/html")
	http.ServeContent(w, r, "", time.Now(), file)
}

func application() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("application.html")

		if err != nil {
			http.Error(w, "500", http.StatusInternalServerError)
			return
		}

		defer file.Close()

		w.Header().Set("Content-Type", "text/html")
		http.ServeContent(w, r, "", time.Now(), file)
	}
}

func main() {
	http.HandleFunc("/", void)
	http.HandleFunc("/login", login)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/application", authentication(application()))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
