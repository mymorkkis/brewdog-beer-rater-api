package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mymorkkis/brewdog-beer-rater-api/db"
)

func RatingsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ratings, err := db.QueryRows(
		"select beer_id, rating from beer_ratings where user_id = $1;",
		parseRows,
		vars["userID"],
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&ratings)
	if err != nil {
		log.Printf("Unable to encode queried data into JSON %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func parseRows(rows db.Rows) ([]db.BeerRating, error) {
	ratings := []db.BeerRating{}

	for rows.Next() {
		var br db.BeerRating

		err := rows.Scan(&br.BeerID, &br.Rating)
		if err != nil {
			log.Printf("Unable to parse ratings by user row %v", err)
			return nil, err
		}

		ratings = append(ratings, br)
	}

	return ratings, nil
}
