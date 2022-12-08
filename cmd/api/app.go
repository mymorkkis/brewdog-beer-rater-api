package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO Inject the router as a dependency?
type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DBPool   *pgxpool.Pool
}

func (app *Application) Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/rating/create", app.RatingCreate).Methods(http.MethodPost)
	r.HandleFunc("/rating/{ratingID:[0-9]+}", app.RatingGet).Methods(http.MethodGet)
	r.HandleFunc("/ratings/{userID:[0-9]+}", app.RatingsByUser).Methods(http.MethodGet)

	r.Use(app.recoverPanic, app.logRequest, secureHeaders)

	return r
}

func (app *Application) serveJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		app.serverError(w, err)
	}
}
