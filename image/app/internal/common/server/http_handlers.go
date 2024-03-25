package server

import (
	"api/internal/common/app"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strings"
)

var store = sessions.NewCookieStore([]byte("SESSION_KEY"))

type HttpHandler struct {
	Application *app.Application
}

func NewHttpHandler(app *app.Application) *HttpHandler {
	return &HttpHandler{
		Application: app,
	}
}

func (h *HttpHandler) LoginForm(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./templates/auth/signin.html", false)
}

func (h *HttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGIN", r.FormValue("email"), r.FormValue("password"))
	result := h.Application.AuthModule.Commands.CheckPassword(r.FormValue("email"), r.FormValue("password"))
	if result == true {
		session, _ := store.Get(r, "session-name")
		session.Values["foo"] = "bar"
		session.Values[42] = 43
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	} else {
		fmt.Println("WRONG")
	}
	RenderTemplate(w, "./templates/auth/signin.html", false)
}

func (h *HttpHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["foo"] = "logout"
	err := session.Save(r, w)
	if err != nil {
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)

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
func RenderTemplate(w http.ResponseWriter, tmpl string, includeLayout bool) {
	if includeLayout == true {
		parsedTemplate, _ := template.ParseFiles("./templates/layout.html")

		err := parsedTemplate.Execute(w, nil)

		if err != nil {
			fmt.Println("Error executing template: ", tmpl, "error: ", err)
		}
	} else {
		parsedTemplate, _ := template.ParseFiles(tmpl)
		err := parsedTemplate.Execute(w, nil)

		if err != nil {
			fmt.Println("Error executing template: ", tmpl, "error: ", err)
		}
	}

}
