package main

import (
	"log"

	"github.com/rileyarnie/ecomm/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to database")
}
