package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mymorkkis/brewdog-beer-rater-api/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ratings/{userID:[0-9]+}", handlers.RatingsByUser).Methods(http.MethodGet)
	http.Handle("/", r)

	log.Println("listening for requests at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
