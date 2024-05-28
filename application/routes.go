package application

import (
	"app-structure/handler"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *App) loadRoutes() *chi.Mux {

	fs := http.FileServer(http.Dir("static"))

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Handle("/static/*", http.StripPrefix("/static/", fs))
	router.Route("/", a.homeRoute)
	a.router = router

	return router
}

func (a *App) homeRoute(router chi.Router) {

	homeHandler := handler.NewHome(a.templates)
	// homeHandler := handler.NewHome()

	router.Get("/", homeHandler.HomePage)

}
