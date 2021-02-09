package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"path"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	cookieName       = "session_example"
	cookieSecureBool = false
	cookieMaxAge     = 600 // Seconds.
	serverTimeout    = 15  // Seconds.
)

type user struct {
	password string
	cookie   string
	stamp    time.Time // Golang's zero date is '0001-01-01 00:00:00 +0000 UTC'.
}

type content struct {
	Title string
}

var (
	users = map[string]*user{}
	mutex sync.Mutex
)

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

		mutex.Lock() // By the way, it is a run-time error to unlock the mutex twice.

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
					MaxAge:   cookieMaxAge,
				})

				user.stamp = time.Now()

				mutex.Unlock()

				h(w, r)
				return
			}
		}

		mutex.Unlock()

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

func main() {
	users["john"] = &user{
		password: "$2a$04$zS2sTndQnwWe53Vy6eQ2NuwvL06sVGDcdwAaRdta4GPbBzE73ZZUK", // smith
		cookie:   "",
	}

	http.HandleFunc("/", staticFiles)
	http.HandleFunc("/login", login)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/session/application", authentication(application()))

	tick := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {

			case <-tick.C:
				fmt.Printf("Cleaning... ")
				mutex.Lock()

				for username, user := range users {
					if user.cookie != "" && time.Now().Sub(user.stamp).Seconds() > serverTimeout {

						fmt.Printf("[timeout user '%s'] ", username)
						user.cookie = ""

					}
				}

				mutex.Unlock()
				fmt.Printf("ok\n")

			case <-quit:
				return
			}
		}
	}()

	server := &http.Server{
		Addr:           "Localhost:8000",
		Handler:        nil, // The DefaultServeMux is used.
		WriteTimeout:   time.Second * 10,
		ReadTimeout:    time.Second * 10,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 32768,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	fmt.Printf("Listening 'Localhost:8000' \n\n")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	tick.Stop()
	close(quit)

	fmt.Printf("\nShutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 10))
	defer cancel()

	server.Shutdown(ctx)

	fmt.Printf(" ok\n")
}
