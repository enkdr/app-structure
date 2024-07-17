package application

import (
	"app-structure/database"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type App struct {
	router    http.Handler // chi
	templates *template.Template
	db        *sql.DB
}

func NewApp() *App {

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	var templatesPath string
	if os.Getenv("DOCKER_ENV") == "true" {
		templatesPath = "templates/index.html"
	} else {
		templatesPath = "../templates/index.html"
	}

	app := &App{
		templates: template.Must(template.ParseFiles(templatesPath)),
		db:        db,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {

	defer a.db.Close()

	server := &http.Server{
		Addr:    ":8001",
		Handler: a.router,
	}

	fmt.Printf("::: starting server on port %s:::", server.Addr)

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

}
