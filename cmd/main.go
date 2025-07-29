package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/liam/pkg/drivers"
	"github.com/schlucht/liam/pkg/models"
)

var session *scs.SessionManager

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	Session       *scs.SessionManager
	DB            models.DBModel
}

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"messge"`
	Data    interface{} `json:"data"`
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
	var cgf config
	tc := make(map[string]*template.Template)
	cgf.port = 5100
	cgf.env = "development"
	cgf.db.dsn = "ots:goweb@/goweb?parseTime=true"

	infoLog := log.New(os.Stdout, "\x1b[32mINFO:\x1b[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "\x1b[31mERROR:\x1b[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	
	var conn *sql.DB
	var err error

	conn, err = drivers.MySqlDB(cgf.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}	
	defer conn.Close()

	mod := models.NewDBModel(conn, infoLog, errorLog)
	app := &application{
		config:        cgf,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       "0.0.1",
		Session:       session,
		DB:            *mod,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		errorLog.Fatal(err)
	}
	
}
