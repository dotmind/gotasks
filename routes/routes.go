package routes

import (
	"encoding/json"
	"net/http"

	"../config"
	wifiname "github.com/yelinaung/wifi-name"
)

type Response struct {
	Success  bool   `json:"success"`
	WifiName string `json:"wifi"`
}

var appConfig = config.Load()

func isOwnNetwork() bool {
	return appConfig.WifiName == wifiname.WifiName()
}

func jsonResponse(w http.ResponseWriter) {
	success := isOwnNetwork()

	response := Response{
		Success:  success,
		WifiName: wifiname.WifiName(),
	}

	json.NewEncoder(w).Encode(response)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w)
}
