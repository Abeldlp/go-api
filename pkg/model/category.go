package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

func GetAllCategories() []Category {
	var categories []Category
	db.Find(&categories)
	return categories
}

func GetCategoryById(Id int64) (*Category, *gorm.DB) {
	var category Category
	db := db.Where("Id=?", Id).Find(&category)
	return &category, db
}

func (c *Category) CreateCategory() *Category {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func DeleteCategory(Id int64) {
	var category Category
	db.Where("Id=?", Id).Delete(category)
}
