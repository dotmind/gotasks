package main

import (
	"log"
	"net/http"

	"./config"
	"./routes"
	"github.com/gorilla/mux"
)

var appConfig = config.Load()

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.RootHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(appConfig.Port, r))
}
