package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DNS = "host=localhost user=diegoalfonsoc18  password=Shadow1827$$ dbname=gorm port=5432 sslmode=disable"
var DB *gorm.DB


func DBConnection () {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}