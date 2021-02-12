package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/templates"
	"github.com/jameycribbs/mule/pkg/models"
)

func ShowExpense(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			app.NotFound(w)
			return
		}

		e, err := app.Models.Expenses.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.NotFound(w)
			} else {
				app.ServerError(w, err)
			}
			return
		}

		data := &templates.TemplateData{Expense: e}

		render(app, w, r, "expenses/show.page.tmpl", data)
	}
}

func CreateExpense(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.ClientError(w, http.StatusMethodNotAllowed)
			return
		}

		// Create some variables holding dummy data. We'll remove these later on
		// during the build.
		name := "Water and Sewer"
		date := time.Now()
		amount := 7515
		category := "Utilities"
		notes := ""

		id, err := app.Models.Expenses.Insert(name, date, amount, category, notes)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/expense?id=%d", id), http.StatusSeeOther)
	}
}
