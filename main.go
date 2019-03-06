package main

import (
	"log"
	"net/http"

	"./config"
	db "./database"
	"./routes"
	"github.com/gorilla/mux"
)

var appConfig = config.Load()

func main() {
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", routes.RootHandler).Methods("GET")
	r.HandleFunc("/auth", routes.AuthHandler).Methods("GET")
	r.HandleFunc("/getall", routes.GetAllHandler).Methods("GET")
	r.HandleFunc("/add", routes.AddHandler).Methods("GET", "POST") // @TODO Remove GET method
	r.HandleFunc("/delete", routes.DeleteHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(appConfig.Port, r))
}
