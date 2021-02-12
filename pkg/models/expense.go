package models

import (
	"time"

	"github.com/jameycribbs/hare"
)

type Expense struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Amount   int       `json:"amount"`
	Category string    `json:"category"`
	Notes    string    `json:"notes"`
}

func (e *Expense) GetID() int {
	return e.ID
}

func (e *Expense) SetID(id int) {
	e.ID = id
}

func (e *Expense) AfterFind() {
	*e = Expense(*e)
}

func QueryExpenses(db *hare.Database, queryFn func(entry Expense) bool, limit int) ([]*Expense, error) {
	var results []*Expense
	var err error

	ids, err := db.IDs("expenses")
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		expense := Expense{}

		if err = db.Find("expenses", id, &expense); err != nil {
			return nil, err
		}

		if queryFn(expense) {
			results = append(results, &expense)
		}

		if limit != 0 && limit == len(results) {
			break
		}
	}

	return results, err
}
