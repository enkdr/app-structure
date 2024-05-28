package application

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var db *sql.DB

type App struct {
	router    http.Handler
	templates *template.Template
	db        *sql.DB
}

func New() *App {

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		fmt.Println("postgres connected")
	}

	// Determine the base path for the templates directory
	var templatesPath string
	if os.Getenv("DOCKER_ENV") == "true" {
		templatesPath = "templates/index.html" // Update with the correct path in your Docker container
	} else {
		templatesPath = "../templates/index.html" // Update with the correct relative path in your local environment
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
			ch <- fmt.Errorf("::: failed to start server: %w :::", err)
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

	return nil
}
