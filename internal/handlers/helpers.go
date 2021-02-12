package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/templates"
)

func addDefaultTemplateData(td *templates.TemplateData, r *http.Request) *templates.TemplateData {
	if td == nil {
		td = &templates.TemplateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}

func render(app *application.Application, w http.ResponseWriter, r *http.Request, name string, td *templates.TemplateData) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	// Write the template to the buffer, instead of straight to the
	// http.ResponseWriter, so we can catch and properly report
	// errors to the user.
	err := ts.Execute(buf, addDefaultTemplateData(td, r))
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}
