package routes

import (
	"YesFix/handlers"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {

	r.Get("/", handlers.HomeHandler)
	// r.Get("/locality", handlers.LocalityHandler)

	// submitHandler := handlers.NewSubmitHandler(application.Query)
	// r.Post("/prebook", submitHandler.ServeHTTP)

	r.Get("/about", handlers.AboutHandler)
	r.NotFound(handlers.NotFoundHandler)
}
