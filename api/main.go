package main

import (
	"log"
	"net/http"

	"./config"
	db "./database"
	router "./routes"
)

var appConfig = config.Load()

func main() {
	db.Init()

	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(appConfig.Port, r))
}
