package handlers

import (
	"fmt"
	"net/http"

	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/templates"
)

func render(app *application.Application, w http.ResponseWriter, r *http.Request, name string, td *templates.TemplateData) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
		app.ServerError(w, err)
	}
}
