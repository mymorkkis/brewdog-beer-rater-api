package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mymorkkis/brewdog-beer-rater-api/handlers"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ratings/{userID:[0-9]+}", handlers.RatingsByUser).Methods(http.MethodGet)
	http.Handle("/", r)

	internal.InfoLog.Println("listening for requests at port 8080...")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
