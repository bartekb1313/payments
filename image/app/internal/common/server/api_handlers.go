package server

import (
	"api/internal/common/app"
	"api/internal/common/server/spec"
	"api/internal/organization/ports"
	"context"
	"fmt"
)

type Server struct {
	Application *app.Application
}

func NewServer(app *app.Application) *Server {
	return &Server{
		Application: app,
	}
}

func (d *Server) PostApiBranches(ctx context.Context, request spec.PostApiBranchesRequestObject) (spec.PostApiBranchesResponseObject, error) {
	handlers := ports.NewHandler(d.Application)
	return handlers.AddBranch(request)
}

func (d *Server) GetApiBranches(ctx context.Context, request spec.GetApiBranchesRequestObject) (spec.GetApiBranchesResponseObject, error) {
	fmt.Println("GET")
	handlers := ports.NewHandler(d.Application)
	return handlers.GetBranches(request)
}
