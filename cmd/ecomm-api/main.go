package main

import (
	"log"

	"github.com/ianschenck/envflag"
	"github.com/rileyarnie/ecomm/db"
	"github.com/rileyarnie/ecomm/ecomm-api/handler"
	"github.com/rileyarnie/ecomm/ecomm-api/server"
	"github.com/rileyarnie/ecomm/ecomm-api/storer"
)

const minSecretKeyLength = 32

func main() {
	var secretKey = envflag.String("SECRET_KEY", "thisismysecretkeythatiwillusetocreatetokens", "secret jey for JWT signing")

	if len(*secretKey) < minSecretKeyLength {
		log.Fatalf("SECRET_KEY must be at least %d characters", minSecretKeyLength)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())

	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
