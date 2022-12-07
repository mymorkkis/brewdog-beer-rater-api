package main

import (
	"net/http"

	"github.com/mymorkkis/brewdog-beer-rater-api/cmd/api"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal/db"
)

func main() {
	// TODO Add router here?
	app := &api.Application{
		ErrorLog: internal.ErrorLog,
		InfoLog:  internal.InfoLog,
	}

	DBPool, err := db.OpenPGXPool()
	if err != nil {
		app.ErrorLog.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer DBPool.Close()

	app.DBPool = DBPool

	app.InfoLog.Println("listening for requests on port 8080...")
	err = http.ListenAndServe(":8080", app.Routes())
	app.ErrorLog.Fatal(err)
}
