package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal/db/services"
)

// TODO Ideally I'd have a ratings sub folder but then it can't find Application,
// would also have the same issue if I create a ratings package.
// Probably need to refactor how this is implemented, maybe pass handlers to Application?

func (app *Application) RatingCreate(w http.ResponseWriter, r *http.Request) {
	service := &services.BeerRatingService{DBPool: app.DBPool}
	// TODO Handle unique constraint with user/beer here

	// TODO Helpful error if these are not correct types
	rating, err := service.InsertRating(
		r.FormValue("userID"),
		r.FormValue("beerID"),
		r.FormValue("rating"),
	)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.serveJSON(w, rating)
}

func (app *Application) RatingGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	service := &services.BeerRatingService{DBPool: app.DBPool}
	rating, err := service.GetRating(vars["ratingID"])
	if err != nil {
		if errors.Is(err, services.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.serveJSON(w, rating)
}

func (app *Application) RatingsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	service := &services.BeerRatingService{DBPool: app.DBPool}
	ratings, err := service.GetRatingsByUser(vars["userID"])
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.serveJSON(w, ratings)
}
