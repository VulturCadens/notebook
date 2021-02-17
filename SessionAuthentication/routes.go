package main

import (
	"net/http"
	"time"
)

type server struct {
	router     *http.ServeMux
	httpServer *http.Server
}

func newServer(addr string) server {
	s := server{}

	s.router = http.NewServeMux()

	s.router.HandleFunc("/", staticFiles)
	s.router.HandleFunc("/login", login)
	s.router.HandleFunc("/welcome", welcome)
	s.router.HandleFunc("/session/application", authentication(application()))

	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        s.router,
		WriteTimeout:   time.Second * 10,
		ReadTimeout:    time.Second * 10,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 32768,
	}

	return s
}
