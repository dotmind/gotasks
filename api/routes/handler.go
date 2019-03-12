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

	UpdateResponse struct {
		Task db.Task `json:"task"`
	}
)

func isOwnNetwork() bool {
	return appConfig.WifiName == wifiname.WifiName()
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	Response{}.WithSuccess().Send(w)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	Response{}.WithSuccess().Send(w)
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tasks := db.GetAllTasks()

	Response{}.WithSuccess().WithPayload(GetAllResponse{
		Total: len(tasks),
		Tasks: tasks,
	}).Send(w)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var task db.Task
	err := decoder.Decode(&task)

	if err != nil {
		panic(err)
	}

	if len(task.Name) == 0 {
		Response{}.WithError().Send(w)
		return
	}

	_, id := db.NewTask(task)
	Response{}.WithSuccess().WithPayload(AddResponse{
		Id: id,
	}).Send(w)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var body db.Task
	err := decoder.Decode(&body)

	if err != nil {
		panic(err)
	}

	_, task, err := db.GetTask(body.Id)

	if err != nil {
		Response{}.WithError().Send(w)
		return
	}

	task.Description = body.Description
	task.Active = body.Active
	db.SaveTask(task)

	Response{}.WithSuccess().WithPayload(UpdateResponse{
		Task: task,
	}).Send(w)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	Response{}.WithSuccess().Send(w)
}
