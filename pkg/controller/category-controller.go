package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Abeldlp/bookManagement/pkg/model"
	"github.com/gorilla/mux"
)

var NewCategory model.Category

func GetCategories(w http.ResponseWriter, r *http.Request) {
	newCategories := model.GetAllCategories()
	res, _ := json.Marshal(newCategories)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	Id, err := strconv.ParseInt(categoryId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	categoryDetails, _ := model.GetCategoryById(Id)
	res, _ := json.Marshal(categoryDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {

}
