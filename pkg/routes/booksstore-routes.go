package routes

import (
	"github.com/Abeldlp/bookManagement/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controller.DeleteBook).Methods("DELETE")
}
