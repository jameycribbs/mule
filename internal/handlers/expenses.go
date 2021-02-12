package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/templates"
	"github.com/jameycribbs/mule/pkg/models"
)

func ShowExpense(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
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

		render(app, w, r, "expenses/show.page.tmpl", &templates.TemplateData{Expense: e})
	}
}

func CreateExpenseForm(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(app, w, r, "expenses/create.page.tmpl", nil)
	}
}

func CreateExpense(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errors := make(map[string]string)

		err := r.ParseForm()
		if err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		name := r.PostForm.Get("name")
		if strings.TrimSpace(name) == "" {
			errors["name"] = "This field cannot be blank"
		}

		date, err := time.Parse("January 2, 2006", r.PostForm.Get("date"))
		if err != nil {
			errors["date"] = "This is not a valid date"
		}

		amount, err := strconv.ParseFloat(r.PostForm.Get("amount"), 64)
		if err != nil {
			errors["amount"] = "This field must be a dollar amount"
		}

		category := r.PostForm.Get("category")
		if strings.TrimSpace(category) == "" {
			errors["category"] = "This field cannot be blank"
		}

		notes := r.PostForm.Get("notes")

		if len(errors) > 0 {
			fmt.Fprint(w, errors)
			return
		}

		id, err := app.Models.Expenses.Insert(name, date, amount, category, notes)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/expense/%d", id), http.StatusSeeOther)
	}
}
