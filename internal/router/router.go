package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/handlers"
)

func New(app *application.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handlers.Home(app))
	r.Get("/expense/create", handlers.CreateExpenseForm(app))
	r.Post("/expense/create", handlers.CreateExpense(app))
	r.Get("/expense/{id}", handlers.ShowExpense(app))

	r.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return handlers.RecoverPanic(app, handlers.LogRequest(app, handlers.SecureHeaders(r)))
}
