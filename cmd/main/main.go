package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

// TODO: Application struct needs a few more properties
// TODO: 1) DB 2) Session manager scs
type application struct {
	config           config
	infoLog          *log.Logger
	errorLog         *log.Logger
	templateCache    map[string]*template.Template
	templateMarkdown map[string]*template.FuncMap
	version          string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Println(fmt.Sprintf("Starting Front end in %s mode on port %d", app.config.env, app.config.port))
	return srv.ListenAndServe()
}

// main driver
func main() {
	//gob.Register(TransactionData{})
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api")
	// TODO read in the connection string for the db
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// This is the DB Driver logic
	//conn, err := driver.OpenDB(cfg.db.dsn)
	//if err != nil {
	//	errorLog.Println(err)
	//}
	//
	//defer conn.Close()

	// setup session
	//session = scs.New()
	//// setup a lifetime
	//session.Lifetime = 24 * time.Hour

	tc := make(map[string]*template.Template)
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
