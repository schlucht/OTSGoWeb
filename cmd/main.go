package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	Session       *scs.SessionManager
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", 5100),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting server on port %s", srv.Addr)
	return srv.ListenAndServe()
}

func main() {

	tc := make(map[string]*template.Template)

	infoLog := log.New(os.Stdout, "\x1b[32mINFO:\x1b[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "\x1b[31mERROR:\x1b[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		Session:       session,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		errorLog.Fatal(err)
	}
}
