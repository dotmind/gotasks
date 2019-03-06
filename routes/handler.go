package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "../database"
	wifiname "github.com/yelinaung/wifi-name"
)

func isOwnNetwork() bool {
	return appConfig.WifiName == wifiname.WifiName()
}

func jsonResponse(w http.ResponseWriter, extra ExtraResponse) {
	success := isOwnNetwork()

	response := Response{
		Success:  success,
		WifiName: wifiname.WifiName(),
		Payload:  extra,
	}

	json.NewEncoder(w).Encode(response)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, ExtraResponse{})
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, ExtraResponse{})
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tasks := db.GetAllTasks()

	fmt.Printf("It's a fish! %#v\n", tasks)
	jsonResponse(w, ExtraResponse{})
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	task := db.Task{
		Name:        "name",
		Description: "description",
		Active:      true,
		Time:        1,
	}

	_, id := db.SaveTask(task)
	jsonResponse(w, ExtraResponse{
		Id: id,
	})
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, ExtraResponse{})
}
