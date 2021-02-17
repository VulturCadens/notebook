package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type content struct {
	Title string
}

var templates = template.Must(template.ParseFiles(
	"./templates/partial/head.html",
	"./templates/application.html",
	"./templates/login.html",
	"./templates/welcome.html",
))

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

	mutex.Lock()

	if _, ok := users[username]; ok {
		err := bcrypt.CompareHashAndPassword([]byte(users[username].password), []byte(password))

		if err != nil {
			mutex.Unlock()

			http.Error(w, "403", http.StatusForbidden)

			return
		}

		/*
		 *  Create a new random slice (byte length of 30 bytes).
		 */

		r := make([]byte, 30)

		if _, err = rand.Read(r); err != nil {
			panic(err)
		}

		cookieValue := fmt.Sprintf("%X", r)

		users[username].cookie = cookieValue
		users[username].stamp = time.Now()

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

			MaxAge: cookieMaxAge,
		})

		fmt.Printf("Username: %s \nPassword: %s \nCookie: %s \nStamp: %v \n\n",
			username,
			users[username].password,
			users[username].cookie,
			users[username].stamp,
		)

		mutex.Unlock()

		content := &content{
			Title: "Welcome",
		}

		if err := templates.ExecuteTemplate(w, "WELCOME", content); err != nil {
			panic(err)
		}

		return
	}

	mutex.Unlock()

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
