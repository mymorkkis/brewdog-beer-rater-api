package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		s := fmt.Sprintf("Hello %v\n", string(b))
		log.Printf(s)
		io.WriteString(w, s)
	}

	http.HandleFunc("/hello", helloHandler)

	log.Println("listening for requests at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
