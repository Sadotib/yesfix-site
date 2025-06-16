package handlers

import (
	notfound "YesFix/views/notFound"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := notfound.NotFound().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}

}
