package main

import (
	"github.com/Abeldlp/bookManagement/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	routes.RegisterCategoryRoutes(r)
	http.Handle("/", r)
	log.Fatal((http.ListenAndServe("localhost:8080", r)))
}
