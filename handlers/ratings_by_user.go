package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/mymorkkis/brewdog-beer-rater-api/db"
)

func RatingsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ratings := db.QueryRows(
		"select beer_id, rating from beer_ratings where user_id = $1;",
		parseRows,
		vars["userID"],
	)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(&ratings)
	if err != nil {
		log.Fatalf("Unable to encode queried data into JSON %v", err)
	}

}

func parseRows(rows pgx.Rows) []db.BeerRating {
	ratings := []db.BeerRating{}

	for rows.Next() {
		var br db.BeerRating

		err := rows.Scan(&br.BeerID, &br.Rating)
		if err != nil {
			log.Fatalf("Unable to pass ratings row %v", err)
		}

		ratings = append(ratings, br)
	}

	return ratings
}
