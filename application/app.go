package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	router http.Handler
}

func NewConstructor() *App {
	app := &App{
		router: LoadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	port := os.Getenv("port")
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           a.router,
		ReadHeaderTimeout: time.Hour * 1,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
	return nil
}
