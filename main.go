package main

import (
	"YesFix/initializers"
	"YesFix/routes"
	"YesFix/types"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

var application *types.App

func init() {
	initializers.LoadEnv()
	query, err := initializers.ConnectToDB()
	if err != nil {
		slog.Error("Failed to connect to the database", "error", err)
		log.Fatal(err)
	}

	application = &types.App{
		Query: query,
	}
}

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")

	slog.Info("HTTP Server started", "listenAddress", listenAddr)

	r := chi.NewMux()

	r.Handle("/public/*", publicdev())
	// r.Get("/*", http.HandlerFunc(publicprod().ServeHTTP))
	// r.Get("/public/*", http.HandlerFunc(publicprod().ServeHTTP))

	routes.Routes(r, application)

	log.Fatal(http.ListenAndServe(listenAddr, r))

}

func publicdev() http.Handler {
	fmt.Println("building static files for development")

	fileSystem := os.DirFS("public")
	fileServer := http.StripPrefix("/public/", http.FileServerFS(fileSystem))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if file exists in "public" directory
		path := strings.TrimPrefix(r.URL.Path, "/public/")
		if _, err := fileSystem.Open(path); err != nil {
			// File not found, trigger 404
			http.NotFound(w, r)
			return
		}

		// File exists, serve it
		fileServer.ServeHTTP(w, r)
	})
}
