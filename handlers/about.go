package handlers

import (
	"YesFix/utils"
	about_page "YesFix/views/about"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, about_page.AboutPage())

}
