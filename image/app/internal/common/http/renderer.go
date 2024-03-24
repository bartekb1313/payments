package http

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	Data      map[string]interface{}
	CSRFToken string
}

func RenderTemplate(w http.ResponseWriter, tmpl string, includeLayout bool, pageContent *TemplateData) {
	if includeLayout == true {
		content, _ := template.ParseFiles("./views/layout.html", "./views/partials/header.html", "./views/partials/sidebar.html", "./views/partials/footer.html", tmpl)
		w.Header().Set("Content-Type", "text/html")
		err := content.Execute(w, pageContent)

		if err != nil {
			fmt.Println("Error executing template: ", tmpl, "error: ", err)
		}
	} else {
		content, _ := template.New("content").ParseFiles(tmpl)
		content.New("base").Parse(`{{ template "content" .}}`)
		err := content.Execute(w, pageContent)

		if err != nil {
			fmt.Println("Error executing template: ", tmpl, "error: ", err)
		}
	}

}
