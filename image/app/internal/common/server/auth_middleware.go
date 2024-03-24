package server

import (
	"fmt"
	"net/http"
)

func AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			session, _ := store.Get(r, "session-name")
			if r.URL.Path != "/login" {

				if session.Values["foo"] != "bar" {
					fmt.Println("REDIRECT")
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}

				fmt.Println("Middleware")
			} else {

				if session.Values["foo"] == "bar" {
					fmt.Println("REDIRECT")
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
			}
			next.ServeHTTP(w, r)
		})

	}
}
