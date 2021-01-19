package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const cookieName = "session_example"

type user struct {
	password string
	cookie   string
}

var users = map[string]*user{}

func authentication(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(w, "401", http.StatusUnauthorized)
				return
			}

			http.Error(w, "400", http.StatusBadRequest)
			return
		}

		for _, user := range users {
			if cookie.Value == user.cookie {
				h(w, r)
				return
			}
		}

		http.Error(w, "403", http.StatusForbidden)
	}
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
		err := bcrypt.CompareHashAndPassword([]byte(users[username].password), []byte(password))

		if err != nil {
			http.Error(w, "401", http.StatusUnauthorized)
			return
		}

		file, err := os.Open("welcome.html")

		if err != nil {
			http.Error(w, "500", http.StatusInternalServerError)
			return
		}

		defer file.Close()

		cookieValue := uuid.New().String()
		users[username].cookie = cookieValue

		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    cookieValue,
			Path:     "/session",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   120,
			// Secure:   true,
			// Expires:  time.Now().Add(120 * time.Second),
		})

		fmt.Printf("Username: %s \nPassword: %s \nCookie: %s \n\n", username, users[username].password, users[username].cookie)

		w.Header().Set("Content-Type", "text/html")
		http.ServeContent(w, r, "", time.Now(), file)
		return
	}

	http.Error(w, "401", http.StatusUnauthorized)
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

func staticFiles(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path

	if p == "/" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, path.Join("./static", p))
}

func main() {
	users["john"] = &user{
		password: "$2a$04$zS2sTndQnwWe53Vy6eQ2NuwvL06sVGDcdwAaRdta4GPbBzE73ZZUK",
		cookie:   "",
	}

	http.HandleFunc("/", staticFiles)
	http.HandleFunc("/login", login)
	http.HandleFunc("/session/welcome", welcome)
	http.HandleFunc("/session/application", authentication(application()))

	fmt.Printf("Listening :8000\n\n")

	http.ListenAndServe(":8000", nil)
}
