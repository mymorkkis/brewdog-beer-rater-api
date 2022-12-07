package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mymorkkis/brewdog-beer-rater-api/db"
)

func (app *Application) RatingsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ratings, err := db.QueryRows(
		"select beer_id, rating from beer_ratings where user_id = $1;",
		app.parseRows,
		vars["userID"],
	)
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

func (app *Application) parseRows(rows db.Rows) ([]db.BeerRating, error) {
	ratings := []db.BeerRating{}

	for rows.Next() {
		var br db.BeerRating

		err := rows.Scan(&br.BeerID, &br.Rating)
		if err != nil {
			app.ErrorLog.Printf("Unable to parse ratings by user row %v", err)
			return nil, err
		}

		ratings = append(ratings, br)
	}

	return ratings, nil
}
