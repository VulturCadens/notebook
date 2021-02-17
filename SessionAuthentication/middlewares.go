package main

import (
	"net/http"
	"time"
)

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
