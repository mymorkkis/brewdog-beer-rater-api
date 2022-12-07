package main

import (
	"log"
	"net/http"

	"github.com/mymorkkis/brewdog-beer-rater-api/cmd/api"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal"
)

func main() {
	app := &api.Application{
		ErrorLog: internal.ErrorLog,
		InfoLog:  internal.InfoLog,
	}

	internal.InfoLog.Println("listening for requests at port 8080...")
	err := http.ListenAndServe(":8080", app.Routes())
	log.Fatal(err)
}
