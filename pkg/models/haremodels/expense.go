package haremodels

import (
	"errors"
	"time"

	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/dberr"
	"github.com/jameycribbs/mule/pkg/models"
)

type ExpenseModel struct {
	DB *hare.Database
}

func (m *ExpenseModel) Insert(name string, date time.Time, amount int, category string, notes string) (int, error) {
	e := models.Expense{Name: name, Date: date, Amount: amount, Category: category, Notes: notes}

	id, err := m.DB.Insert("expenses", &e)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ExpenseModel) Get(id int) (*models.Expense, error) {
	e := &models.Expense{}

	err := m.DB.Find("expenses", id, e)
	if err != nil {
		if errors.Is(err, dberr.NoRecord) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (m *ExpenseModel) Latest(daysAgo time.Duration) ([]*models.Expense, error) {
	now := time.Now()
	timeOfDaysAgo := now.Add(-time.Hour * 24 * daysAgo)

	recs, err := models.QueryExpenses(m.DB, func(e models.Expense) bool {
		return e.Date.After(timeOfDaysAgo) && e.Date.Before(now)
	}, 0)
	if err != nil {
		return nil, err
	}

	return recs, nil
}
