package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
	"github.com/jameycribbs/mule/internal/application"
	"github.com/jameycribbs/mule/internal/router"
	"github.com/jameycribbs/mule/internal/templates"
	"github.com/jameycribbs/mule/pkg/models/haremodels"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dbDir := flag.String("dbDir", "./data", "Hare database directory")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dbDir)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := templates.NewTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := application.New(errorLog, infoLog, haremodels.New(db), templateCache)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  router.New(app),
	}

	infoLog.Printf("Starting server on %s", *addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dbDir string) (*hare.Database, error) {
	ds, err := disk.New(dbDir, ".json")
	if err != nil {
		return nil, err
	}

	db, err := hare.New(ds)
	if err != nil {
		return nil, err
	}

	return db, nil
}
