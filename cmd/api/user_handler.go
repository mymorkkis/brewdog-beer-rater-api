package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal/db/services"
)

func (app *Application) UserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	service := &services.UserService{DBPool: app.DBPool}
	user, err := service.Get(vars["userID"])
	if err != nil {
		if errors.Is(err, services.ErrNoRecord) {
			app.clientErrorWithStatusMessage(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.serveJSON(w, user)
}

func (app *Application) UserCreate(w http.ResponseWriter, r *http.Request) {
	service := &services.UserService{DBPool: app.DBPool}

	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.serverError(w, err)
	}

	var user services.User

	err = json.Unmarshal(rBody, &user)
	if err != nil {
		app.serverError(w, err)
	}

	// TODO Validate fields

	stored_user, err := service.Insert(user.Email, user.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.serveJSON(w, stored_user)
}
