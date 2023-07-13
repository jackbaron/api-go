package main

import (
	"log"
	"net/http"

	"github.com/nhatth/api-service/internal/app/database"
	"github.com/nhatth/api-service/internal/app/router"
	"github.com/nhatth/api-service/pkg/utils"
)

func main() {
	//? Init router
	r := router.InitRouter()

	//? Load config file
	config, err := utils.LoadConfig("./../")

	if err != nil {

		log.Println("Cannot load config file")

		return
	}

	//? Connection DB

	db := database.ConnectDatabase(config)

	log.Println(db)

	err = http.ListenAndServe(":8000", r)

	if err != nil {
		log.Println("Running server faild")
	}

	log.Println("Running server port 8000")
}
