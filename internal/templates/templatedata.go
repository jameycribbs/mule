package templates

import "github.com/jameycribbs/mule/pkg/models"

type TemplateData struct {
	CurrentYear int
	Expense     *models.Expense
	Expenses    []*models.Expense
}
