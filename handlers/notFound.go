package handlers

import (
	notfound "YesFix/views/notFound"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	if r.Header.Get("HX-Request") == "true" {
		// HTMX request – return a partial (fragment)
		err := notfound.NotFound().Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	} else {
		// Normal full-page request
		err := notfound.NotFound().Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	}
}
