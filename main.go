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
	r.HandleFunc("/auth", routes.AuthHandler).Methods("GET")
	r.HandleFunc("/add", routes.AddHandler).Methods("POST")
	r.HandleFunc("/delete", routes.DeleteHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(appConfig.Port, r))
}
