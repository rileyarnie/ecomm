package main

import (
	"log"

	"github.com/rileyarnie/ecomm/db"
	"github.com/rileyarnie/ecomm/ecomm-api/handler"
	"github.com/rileyarnie/ecomm/ecomm-api/server"
	"github.com/rileyarnie/ecomm/ecomm-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())

	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
