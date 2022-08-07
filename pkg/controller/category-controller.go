package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Abeldlp/bookManagement/pkg/model"
	"github.com/Abeldlp/bookManagement/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		fmt.Println("error while parsing id")
	}

	categoryDetails, _ := model.GetCategoryById(Id)
	res, _ := json.Marshal(categoryDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	requestCategory := &model.Category{}
	utils.ParseBody(r, requestCategory)
	createdCategory := requestCategory.CreateCategory()
	res, _ := json.Marshal(createdCategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	Id, err := strconv.ParseInt(categoryId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing id")
	}

	model.DeleteCategory(Id)
	res, _ := json.Marshal("book deleted")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBooksByCategoryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	Id, err := strconv.ParseInt(categoryId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing id")
	}

	categoryBooks, _ := model.GetBooksOfCategory(Id)
	res, _ := json.Marshal(categoryBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	requestCategory := &model.Category{}
	utils.ParseBody(r, requestCategory)

	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	Id, err := strconv.ParseInt(categoryId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	categoryDetails, db := model.GetCategoryById(Id)

	//TODO validation here

	db.Save(&categoryDetails)
	res, _ := json.Marshal(categoryDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
