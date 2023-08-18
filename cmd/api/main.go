package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"10seconds.co/internal/driver"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *driver.DB
}

func main() {
	var cfg config
	cfg.port = 8081

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := "host=localhost port=5432 user=postgres password=password dbname=seconds sslmode=disable timezone=UTC+3 connect_timeout=5"
	db, err := driver.ConnectPostgress(dsn)
	if err != nil {
		log.Fatal("Can not connect to database")
	}

	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errLog,
		db:       db,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("Api listening on port", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()

}
