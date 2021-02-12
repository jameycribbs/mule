package templates

import "github.com/jameycribbs/mule/pkg/models"

type TemplateData struct {
	Expense  *models.Expense
	Expenses []*models.Expense
}
