package main

import (
	"YesFix/initializers"
	"YesFix/routes"
	"embed"

	"YesFix/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

//go:embed public
var publicFS embed.FS

func init() {
	initializers.LoadEnv()
	// query, err := initializers.ConnectToDB()
	// if err != nil {
	// 	slog.Error("Failed to connect to the database", "error", err)
	// 	log.Fatal(err)
	// }

	// application = &types.App{
	// 	Query: query,
	// }
}

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")

	slog.Info("HTTP Server started", "listenAddress", listenAddr)

	r := chi.NewMux()

	r.Handle("/*", public())
	r.NotFound(handlers.NotFoundHandler)

	// r.Get("/*", http.HandlerFunc(publicprod().ServeHTTP))
	// r.Get("/public/*", http.HandlerFunc(publicprod().ServeHTTP))

	routes.Routes(r)

	log.Fatal(http.ListenAndServe(listenAddr, r))

}

// func publicdev() http.Handler {
// 	fmt.Println("building static files for development")

// 	fileSystem := os.DirFS("public")
// 	fileServer := http.StripPrefix("/public/", http.FileServerFS(fileSystem))

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Check if file exists in "public" directory
// 		path := strings.TrimPrefix(r.URL.Path, "/public/")
// 		if _, err := fileSystem.Open(path); err != nil {
// 			// File not found, trigger 404
// 			http.NotFound(w, r)
// 			return
// 		}

// 		// File exists, serve it
// 		fileServer.ServeHTTP(w, r)
// 	})
// }

// func publicprod() http.Handler {
// 	subFS, _ := fs.Sub(publicFS, "/public") // Embed only "public/" folder
// 	fileServer := http.FileServer(http.FS(subFS))

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		path := strings.TrimPrefix(r.URL.Path, "/")
// 		// Check if file exists
// 		if _, err := subFS.Open(path); err != nil {
// 			// If file doesn't exist, delegate to NotFound
// 			http.NotFound(w, r)
// 			return
// 		}
// 		fileServer.ServeHTTP(w, r)
// 	})

// 	// return http.StripPrefix("/public/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	// 	// Check if file exists
// 	// 	path := r.URL.Path
// 	// 	if _, err := subFS.Open(path); err != nil {
// 	// 		http.NotFound(w, r)
// 	// 		return
// 	// 	}
// 	// 	fileServer.ServeHTTP(w, r)
// 	// }))
// }

func public() http.Handler {
	return http.FileServerFS(publicFS)
}
