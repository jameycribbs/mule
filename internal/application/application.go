package application

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/jameycribbs/mule/pkg/models/haremodels"
)

type Application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Models        *haremodels.Models
	TemplateCache map[string]*template.Template
}

func New(errorLog *log.Logger, infoLog *log.Logger, models *haremodels.Models, templateCache map[string]*template.Template) *Application {
	return &Application{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		Models:        models,
		TemplateCache: templateCache,
	}
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
