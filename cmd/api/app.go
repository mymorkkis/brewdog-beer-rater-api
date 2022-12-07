package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DBPool   *pgxpool.Pool
}

func (app *Application) Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ratings/{userID:[0-9]+}", app.RatingsByUser).Methods(http.MethodGet)
	http.Handle("/", r)

	return r
}
