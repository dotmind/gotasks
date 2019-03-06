package routes

import (
	"encoding/json"
	"net/http"

	db "../database"
	wifiname "github.com/yelinaung/wifi-name"
)

type (
	GetAllResponse struct {
		Total int       `json:"total"`
		Tasks []db.Task `json:"tasks"`
	}

	AddResponse struct {
		Id string `json:"id"`
	}
)

func isOwnNetwork() bool {
	return appConfig.WifiName == wifiname.WifiName()
}

func jsonResponse(w http.ResponseWriter, extra interface{}) {
	success := isOwnNetwork()

	response := Response{
		Success: success,
		Payload: EmptyResponse{},
	}

	if response.Success == true {
		response.Payload = extra
	}

	json.NewEncoder(w).Encode(response)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, EmptyResponse{})
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, EmptyResponse{})
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tasks := db.GetAllTasks()

	jsonResponse(w, GetAllResponse{
		Total: len(tasks),
		Tasks: tasks,
	})
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	task := db.Task{
		Name:        "name",
		Description: "description",
		Active:      true,
		Time:        1,
	}

	_, id := db.SaveTask(task)
	jsonResponse(w, AddResponse{
		Id: id,
	})
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, EmptyResponse{})
}
