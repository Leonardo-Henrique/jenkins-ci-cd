package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Página inicial")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
