package main

import (
	"fmt"
	"go-chi-template/db"
	"go-chi-template/services"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

var port = os.Getenv("PORT")

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	return srv.ListenAndServe()
}

func main() {
	var cfg Config
	cfg.Port = port

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
