package main

import (
	"log"
	"net/http"

	"github.com/RianIhsan/go-book-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBook(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9001", r))
}
