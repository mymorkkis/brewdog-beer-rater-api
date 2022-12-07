package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal/db/services"
)

func (app *Application) RatingsByUser(w http.ResponseWriter, r *http.Request) {
	// TODO Add mux to Application?
	vars := mux.Vars(r)

	service := &services.BeerRatingService{DBPool: app.DBPool}
	ratings, err := service.GetRatingsByUser(vars["userID"])
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&ratings)
	if err != nil {
		app.serverError(w, err)
	}
}
