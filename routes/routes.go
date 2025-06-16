package routes

import (
	"YesFix/handlers"
	"YesFix/types"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux, application *types.App) {

	r.Get("/", handlers.HomeHandler)
	r.Get("/locality", handlers.LocalityHandler)

	submitHandler := handlers.NewSubmitHandler(application.Query)
	r.Post("/prebook", submitHandler.ServeHTTP)

	r.NotFound(handlers.NotFoundHandler)
}
