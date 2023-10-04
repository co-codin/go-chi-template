package main

import (
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
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
}
