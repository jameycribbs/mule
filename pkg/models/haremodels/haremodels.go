package haremodels

import (
	"github.com/jameycribbs/hare"
)

type Models struct {
	Expenses *ExpenseModel
}

func New(db *hare.Database) *Models {
	return &Models{
		Expenses: &ExpenseModel{DB: db},
	}
}
