package handlers

import (
	"YesFix/utils"
	home_page "YesFix/views/home"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, home_page.HomePage())

}
