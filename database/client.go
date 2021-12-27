package database

import (
	"golang-sql-associated/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	Connector, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	log.Println("Connection to database succesfull!")
	return nil
}

func Migrate() {

	Connector.Migrator().CreateTable(&models.MasterEvent{})
	Connector.Migrator().CreateTable(&models.EventOrder{})

	Connector.Migrator().CreateConstraint(&models.MasterEvent{}, "EventOrder")
	Connector.Migrator().CreateConstraint(&models.MasterEvent{}, "fk_event_order")
}
