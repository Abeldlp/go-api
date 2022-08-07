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

var RegisterCategoryRoutes = func(router *mux.Router) {
	router.HandleFunc("/categories", controller.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", controller.GetCategories).Methods("GET")
	router.HandleFunc("/categories/{categoryId}", controller.GetCategoryById).Methods("GET")
	router.HandleFunc("/categories/{categoryId}", controller.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{categoryId}", controller.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/categories/{categoriesId}/books", controller.GetBooksByCategoryId).Methods("GET")
}
