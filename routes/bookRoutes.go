package routes

import (
	"github.com/RianIhsan/go-book-api/controllers"
	"github.com/gorilla/mux"
)

var RegisterBook = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books", controllers.GetBook).Methods("GET")
	router.HandleFunc(".api/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
