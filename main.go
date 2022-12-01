package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mymorkkis/brewdog-beer-rater-api/db"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		s := fmt.Sprintf("Hello %v\n", string(b))
		log.Printf(s)
		io.WriteString(w, s)
	}

	http.HandleFunc("/hello", helloHandler)

	dbpool := db.Connect()
	defer dbpool.Close()
	db.GreetingTest(dbpool)

	log.Println("listening for requests at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
