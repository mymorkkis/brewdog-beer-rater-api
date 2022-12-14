package api

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)
}

func (app *Application) clientErrorWithStatusMessage(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
