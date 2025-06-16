package parseValidate

import (
	"YesFix/views/components"
	"log/slog"
	"net/http"
)

func ParseAndValidate(w http.ResponseWriter, r *http.Request) (components.FormValues, map[string]string) {
	err := r.ParseForm()
	if err != nil {
		slog.Error("Failed to parse form data", "error", err)
	}

	formValues := components.FormValues{
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
		City:      r.FormValue("city"),
		Locality:  r.FormValue("locality"),
	}

	errors := formValues.Validate()

	if len(errors) > 0 {
		slog.Error("Validation errors occurred", "errors", errors)
		// return formValues, errors
		return formValues, errors

	}

	return formValues, nil

	// formMap:= make(map[string]string)
	// for key, values := range r.Form {
	// 	if len(values) > 0 {
	// 		formMap[key] = values[0]
	// 	}
	// }

	// return formMap, nil
}
