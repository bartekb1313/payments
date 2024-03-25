package ports

import (
	"api/internal/common/app"
	"api/internal/common/server/spec"
)

type Handler struct {
	app *app.Application
}

func (h *Handler) AddBranch(request spec.PostApiBranchesRequestObject) (spec.PostApiBranches201Response, error) {
	h.app.OrganizationModule.Commands.CreateBranch(request.Body.Name)
	resp := spec.PostApiBranches201Response{}
	return resp, nil
}

func (h *Handler) GetBranches(request spec.GetApiBranchesRequestObject) (spec.GetApiBranches200JSONResponse, error) {
	branches, _ := h.app.OrganizationModule.Queries.GetBranches()
	resp := spec.GetApiBranches200JSONResponse{}
	for _, branch := range branches {
		resp = append(resp, spec.Branch{
			Name: branch.Name(),
		})
	}
	return resp, nil
}

func NewHandler(app *app.Application) *Handler {
	return &Handler{
		app,
	}
}
