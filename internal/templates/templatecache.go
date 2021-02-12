package templates

import (
	"html/template"
	"os"
	"path/filepath"
)

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		matched, err := filepath.Match("*.page.tmpl", info.Name())
		if err != nil {
			return err
		}

		if matched {

			ts, err := template.New(info.Name()).Funcs(functions).ParseFiles(path)
			if err != nil {
				return err
			}

			//ts, err := template.ParseFiles(path)
			//if err != nil {
			//	return err
			//}

			ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
			if err != nil {
				return err
			}

			ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}

			cache[relPath] = ts
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return cache, nil
}
