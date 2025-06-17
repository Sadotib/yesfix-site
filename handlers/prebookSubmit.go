package handlers

import (
	"YesFix/dbConfig/queries"
	"YesFix/utils"
	"YesFix/utils/generateUsid"
	mailconfig "YesFix/utils/mailConfig"
	"YesFix/utils/parseValidate"
	"YesFix/views/components"
	"errors"
	"log"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

type SubmitHandler struct {
	Queries *queries.Queries
}

func NewSubmitHandler(q *queries.Queries) *SubmitHandler {
	return &SubmitHandler{Queries: q}
}

func (h *SubmitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//parse and validate the form data and return display error messages on form if any
	input, errorss := parseValidate.ParseAndValidate(w, r)
	if len(errorss) > 0 {
		utils.Render(w, r, components.Form(input, errorss))
		return
	}

	//if no errors, proceed with inserting the data into the database
	slog.Info("Form data received", "data", input)

	//generate the usid for each user
	usid := generateUsid.GenerateUsid(input.City)

	values := queries.InsertUserTestParams{
		Usid:      usid,
		Firstname: input.FirstName,
		Lastname:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		City:      input.City,
	}

	// Insert the user data into the database
	err := InsertToDB(h, w, r, values)
	if err != nil {
		if err.Code == "23505" {
			// Unique constraint violation
			input.Duplicate = err.ColumnName
			log.Println(err.SchemaName, err.Severity, err.Code, err.Message, err.ColumnName)
			errorss := map[string]string{
				"duplicate": "*Booking already exists for given Email",
			}
			utils.Render(w, r, components.Form(input, errorss))
			return
		}
	}

	// var input queries.InsertUserTestParams
	// w.Header().Set("Content-Type", "application/json")
	// e := json.NewDecoder(r.Body).Decode(&input)

	// if e != nil && e.Error() == "EOF" {
	// 	slog.Error("No user data provided", "error", e)
	// 	http.Error(w, "No user data provided", http.StatusBadRequest)
	// 	return
	// } else if e != nil {
	// 	slog.Error("Failed to read request body", "error", e)
	// 	http.Error(w, "Failed to read request body", http.StatusBadRequest)
	// 	return
	// }
	// defer r.Body.Close()

	// w.WriteHeader(http.StatusOK)
	// jsonstr, err := json.Marshal(row)
	// if err != nil {
	// 	slog.Error("Failed to marshal users", "error", err)
	// 	http.Error(w, "Failed to process users", http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(jsonstr)
}

func InsertToDB(h *SubmitHandler, w http.ResponseWriter, r *http.Request, values queries.InsertUserTestParams) *pgconn.PgError {
	row, err := h.Queries.InsertUserTest(r.Context(), values)
	if err != nil {
		slog.Error("Failed to insert user", "error", err)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			// Unique constraint violation
			log.Println(pgErr.Code)
			return pgErr
		}

	}
	slog.Info("User data inserted successfully", "data", row)

	go func(email string) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("Panic while sending email", "recover", r)
			}
		}()
		mailconfig.SendPrebookMail(email)
	}(values.Email)

	// mailconfig.SendPrebookMail(values.Email)
	w.WriteHeader(http.StatusOK)
	utils.Render(w, r, components.Thanks())
	return nil
}
