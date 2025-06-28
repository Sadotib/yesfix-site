// go:build !dev
//go:build !dev
// +build !dev

package main

import (
	"io/fs"
	"net/http"
	"strings"
)

// go:embed public

// func publicprod() http.Handler {
// 	return http.FileServerFS(publicFS)
// }

func publicprod() http.Handler {
	subFS, _ := fs.Sub(publicFS, "public") // Embed only "public/" folder
	fileServer := http.FileServer(http.FS(subFS))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		// Check if file exists
		if _, err := subFS.Open(path); err != nil {
			// If file doesn't exist, delegate to NotFound
			http.NotFound(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}
