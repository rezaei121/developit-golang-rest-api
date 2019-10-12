package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Connect() (db *gorm.DB) {
	connection, error := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=ehsan dbname=go-db password=123")
	if error != nil {
		log.Fatal(error)
	}

	return connection
}
