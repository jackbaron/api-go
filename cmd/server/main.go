package main

import (
	"log"
	"net/http"

	"github.com/nhatth/api-service/internal/app/database"
	"github.com/nhatth/api-service/internal/app/helpers"
	"github.com/nhatth/api-service/internal/app/routes"
)

func main() {

	//? Load config file
	config, err := helpers.LoadConfig("./../")

	if err != nil {

		log.Fatalln("Cannot load config file")

		return
	}

	//? Connection DB
	db := database.ConnectDatabase(config)

	r := routes.SetUpRoutes(db)

	err = http.ListenAndServe(":8000", r)

	if err != nil {
		log.Println("Running server faild")
	}

	log.Println("Running server port 8000")
}
