package server

import (
	"api/internal/common/app"
	http_helpers "api/internal/common/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"strings"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type HttpHandler struct {
	Application *app.Application
}

func NewHttpHandler(app *app.Application) *HttpHandler {
	return &HttpHandler{
		Application: app,
	}
}

func (h *HttpHandler) LoginForm(w http.ResponseWriter, r *http.Request) {
	http_helpers.Render(w, "auth", "signin", true, http_helpers.TemplateData{})
}

func (h *HttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	result := h.Application.AuthModule.Commands.CheckPassword(r.FormValue("email"), r.FormValue("password"))
	if result == true {

		session, _ := store.Get(r, "payments-session")
		session.Values["email"] = r.FormValue("email")
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Println("WRONG")
	}

	http_helpers.Render(w, "auth", "signin", false, &http_helpers.TemplateData{})
}

func (h *HttpHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "payments-name")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func (h *HttpHandler) FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
