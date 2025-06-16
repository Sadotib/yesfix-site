package handlers

import (
	"YesFix/utils"
	"YesFix/views/home"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, home.Index())

}
