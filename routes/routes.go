package routes

import (
	"YesFix/handlers"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {

	r.Get("/", handlers.HomeHandler)
	r.Get("/about", handlers.AboutHandler)

}
