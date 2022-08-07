package model

import (
	"github.com/Abeldlp/bookManagement/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(
		&Category{},
		&Book{},
	)
}
