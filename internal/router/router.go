package router

import (
	"net/http"

	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/handlers"
)

func New(app *application.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home(app))
	mux.HandleFunc("/expense", handlers.ShowExpense(app))
	mux.HandleFunc("/expense/create", handlers.CreateExpense(app))

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return mux
}
