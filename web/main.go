package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/themilar/msgsqr/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	messages *models.MessageModel
}
type config struct {
	addr      string
	staticDir string
}

var (
	user     = "postgres"
	password = os.Getenv("POSTGRES_PASSWORD")
	host     = "localhost"
	db_name  = "msgsqr"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db_url := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, password, host, db_name)
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		errorLog.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		messages: &models.MessageModel{DB: db},
	}
	var cfg config

	flag.StringVar(&cfg.addr, "addr", "localhost:4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-directory", "./ui/static", "static file directory")
	flag.Parse()

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", cfg.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
