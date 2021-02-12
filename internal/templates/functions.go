package templates

import (
	"html/template"
	"time"
)

func humanDate(t time.Time) string {
	return t.Format("January 02, 2006")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
