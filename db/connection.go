package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DSN = "host=localhost user=diegoalfonsoc18 password=Shadow1827$$ dbname=gorm port=5432 sslmode=disable"


var DB *gorm.DB


func DBConnection () {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatalf("Failed to connect to database: %v", error)
		
	} else {
		log.Println("DB connected")
	}
}