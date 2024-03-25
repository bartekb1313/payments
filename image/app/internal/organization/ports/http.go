package ports

import (
	"api/internal/common/app"
	http_helpers "api/internal/common/http"
	"net/http"
)

type HttpHandler struct {
	app *app.Application
}

func NewHttpHandler(app *app.Application) *HttpHandler {
	return &HttpHandler{
		app: app,
	}
}

type CreateBranchForm struct {
	Name string `schema:"name,required" validate:"required,min=3,max=20"`
}

func (h *HttpHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	http_helpers.RenderTemplate(w, "./views/pages/dashboard.html", true, &http_helpers.TemplateData{})
}

func (h *HttpHandler) BranchList(w http.ResponseWriter, r *http.Request) {
	includeLayout := r.URL.Query().Has("layout")
	branches := make(map[string]interface{})
	bb, _ := h.app.OrganizationModule.Queries.GetBranches()
	for i := 0; i < len(bb); i++ {
		branches[bb[i].Name()] = bb[i].AsView()
	}
	http_helpers.RenderTemplate(w, "./views/pages/organization/branches/list.html", !includeLayout, &http_helpers.TemplateData{Data: branches})
}

func (h *HttpHandler) BranchForm(w http.ResponseWriter, r *http.Request) {
	includeLayout := r.URL.Query().Has("layout")
	http_helpers.RenderTemplate(w, "./views/pages/organization/branches/create.html", !includeLayout, &http_helpers.TemplateData{})

}

func (h *HttpHandler) BranchCreate(w http.ResponseWriter, r *http.Request) {
	var branchForm = &CreateBranchForm{}
	http_helpers.Populate(branchForm, r)

	validationErrors := http_helpers.Validate(branchForm)

	if len(validationErrors) == 0 {
		h.app.OrganizationModule.Commands.CreateBranch(branchForm.Name)
		http.Redirect(w, r, "/organization/branches/list?layout=false", http.StatusSeeOther)
	} else {
		includeLayout := r.URL.Query().Has("layout")
		http_helpers.RenderTemplate(w, "./views/pages/organization/branches/create.html", !includeLayout, &http_helpers.TemplateData{StringMap: validationErrors})
	}
}
