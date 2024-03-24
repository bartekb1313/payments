package server

import (
	"api/internal/common/app"
	organization_handlers "api/internal/organization/ports"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InitRoutes(app *app.Application) func(r chi.Router) {
	coreHandlers := NewHttpHandler(app)
	organizationHandlers := organization_handlers.NewHttpHandler(app)
	return func(r chi.Router) {
		r.Use(AuthMiddleware())
		coreHandlers.FileServer(r, "/assets", http.Dir("./assets"))
		r.Get("/login", coreHandlers.LoginForm)
		r.Get("/logout", coreHandlers.Logout)
		r.Post("/login", coreHandlers.Login)
		r.Get("/", organizationHandlers.Dashboard)
		r.Get("/blank", organizationHandlers.Blank)
		r.Get("/organization/branches/list", organizationHandlers.BranchList)
		r.Get("/organization/branches/create", organizationHandlers.BranchForm)
		r.Post("/organization/branches/create", organizationHandlers.BranchCreate)

	}
}
