package handlers

import (
	"YesFix/utils/locationData"
	"YesFix/views/components"
	"fmt"
	"net/http"
)

func LocalityHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	fmt.Println("City:", city)
	localities := locationData.Cities[city]
	components.LocalityOptions(localities).Render(r.Context(), w)
}
