package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	cookieName       = "session_example"
	cookieSecureBool = false
	cookieMaxAge     = 600 // Seconds.
	serverTimeout    = 15  // Seconds.
)

// A fake database with some mocking data.

type user struct {
	password string
	cookie   string
	stamp    time.Time // Golang's zero date is '0001-01-01 00:00:00 +0000 UTC'.
}

var (
	users = map[string]*user{}
	mutex sync.Mutex
)

func init() {
	users["john"] = &user{
		password: "$2a$04$zS2sTndQnwWe53Vy6eQ2NuwvL06sVGDcdwAaRdta4GPbBzE73ZZUK", // smith
		cookie:   "",
	}
}

func main() {
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

	server := newServer("Localhost:8000")

	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	fmt.Printf("Listening 'Localhost:8000'\n\n")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	tick.Stop()
	close(quit)

	fmt.Printf("\nShutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 10))
	defer cancel()

	server.httpServer.Shutdown(ctx)

	fmt.Printf("ok\n")
}
