package handlers

import (
	"net/http"

	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/templates"
)

func Home(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e, err := app.Models.Expenses.Latest(30)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		render(app, w, r, "home.page.tmpl", &templates.TemplateData{Expenses: e})
	}
}
