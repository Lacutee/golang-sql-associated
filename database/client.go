package database

import (
	"golang-sql-associated/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection to database succesfull!")
	return nil
}

func Migrate() {
	Connector.AutoMigrate(&models.EventOrder{}, &models.MasterEvent{})
}
