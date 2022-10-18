package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect(connectioString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectioString), &gorm.Config{})
	if dbError != nil {
		log.Fatal("error configuring the database: ", dbError)
	}
	log.Println("Hey! You successfully connected to your CockroachDB cluster.")
}
