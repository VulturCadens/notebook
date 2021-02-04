package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	cookieName       = "session_example"
	cookieSecureBool = false
)

type user struct {
	password string
	cookie   string
	stamp    time.Time // Golang's zero date is '0001-01-01 00:00:00 +0000 UTC'.
}

type content struct {
	Title string
}

var users = map[string]*user{}

var templates = template.Must(template.ParseFiles(
	"./templates/partial/head.html",
	"./templates/application.html",
	"./templates/login.html",
	"./templates/welcome.html",
))

func authentication(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)

		if err != nil {

			/*
			 *  if errors.Is(err, http.ErrNoCookie) {
			 *    http.Error(w, "400", http.StatusBadRequest)
			 *    return
			 *  }
			 */

			http.Error(w, "400", http.StatusBadRequest)
			return
		}

		for _, user := range users {
			if cookie.Value == user.cookie {

				/*
				 *  Refreshing a session token.
				 */

				http.SetCookie(w, &http.Cookie{
					Name:     cookieName,
					Value:    cookie.Value,
					Path:     "/session",
					HttpOnly: true,
					SameSite: http.SameSiteStrictMode,
					Secure:   cookieSecureBool,
					MaxAge:   30,
				})

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
			http.Error(w, "403", http.StatusForbidden)
			return
		}

		users[username].stamp = time.Now()

		cookieValue := uuid.New().String()
		users[username].cookie = cookieValue

		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    cookieValue,
			Path:     "/session",
			HttpOnly: true, // Prevent access to cookie values via JavaScript.
			SameSite: http.SameSiteStrictMode,

			/*
			 *  A cookie is sent to the server only with an encrypted request over the HTTPS protocol.
			 */

			Secure: cookieSecureBool,

			/*
			 *  A session cookie is a cookie without an expiration time.
			 *
			 *  Session cookies are deleted when the browser closes (the session ends) and
			 *  permanent cookies are deleted at a date specified by the 'Expires' attribute or
			 *  after a period of time specified by the 'Max-Age' attribute.
			 *
			 *  Expires:  time.Now().Add(120 * time.Second)
			 */

			MaxAge: 30, // Seconds.
		})

		fmt.Printf("Username: %s \nPassword: %s \nCookie: %s \nStamp: %v \n\n",
			username,
			users[username].password,
			users[username].cookie,
			users[username].stamp,
		)

		content := &content{
			Title: "Welcome",
		}

		if err := templates.ExecuteTemplate(w, "WELCOME", content); err != nil {
			panic(err)
		}

		return
	}

	http.Error(w, "403", http.StatusForbidden)
}

func login(w http.ResponseWriter, r *http.Request) {
	content := &content{
		Title: "Login",
	}

	if err := templates.ExecuteTemplate(w, "LOGIN", content); err != nil {
		panic(err)
	}
}

func application() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := &content{
			Title: "Application",
		}

		if err := templates.ExecuteTemplate(w, "APPLICATION", content); err != nil {
			panic(err)
		}
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
		password: "$2a$04$zS2sTndQnwWe53Vy6eQ2NuwvL06sVGDcdwAaRdta4GPbBzE73ZZUK", // smith
		cookie:   "",
	}

	http.HandleFunc("/", staticFiles)
	http.HandleFunc("/login", login)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/session/application", authentication(application()))

	fmt.Printf("Listening 'Localhost:8000'...\n\n")

	http.ListenAndServe("Localhost:8000", nil)
}
