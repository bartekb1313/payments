package server

import (
	"net/http"
)

func AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			session, _ := store.Get(r, "payments-session")

			if r.URL.Path != "/login" {

				if _, ok := session.Values["email"]; !ok {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
			} else {
				if _, ok := session.Values["email"]; ok {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
			}
			next.ServeHTTP(w, r)
		})

	}
}
