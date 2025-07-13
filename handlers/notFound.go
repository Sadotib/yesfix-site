package handlers

import (
	"YesFix/utils"
	notfound "YesFix/views/notFound"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := utils.Render(w, r, notfound.NotFound())
	if err != nil {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}

}
