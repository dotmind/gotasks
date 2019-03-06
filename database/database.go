package database

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
)

type (
	Task struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Active      bool   `json:"active"`
		Time        int    `json:"time"`
		Username    string `json:"username"`
	}

	Database *scribble.Driver
)

const dir string = "./database/.data"
const taskEntry string = "task"

var db *scribble.Driver

func Init() {
	driver, err := scribble.New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	db = driver
}

func save(entry string, input string, node interface{}) Database {
	db.Write(entry, input, node)
	return db
}

func generateUUID() (string, uuid.UUID) {
	id := uuid.New()
	return id.String(), id
}

func SaveTask(task Task) (Database, string) {
	id, _ := generateUUID()
	task.Id = id
	return save(taskEntry, id, task), id
}

func GetTask(id string) (Database, Task) {
	task := Task{}
	db.Read(taskEntry, id, &task)
	return db, task
}

func GetAllTasks() (Database, []Task) {
	tasksId, _ := db.ReadAll(taskEntry)

	tasks := []Task{}
	for _, task := range tasksId {
		t := Task{}
		json.Unmarshal([]byte(task), &t)
		tasks = append(tasks, t)
	}

	return db, tasks
}

func DeleteTask(id string) Database {
	db.Delete(taskEntry, id)
	return db
}

func ClearAllTasks() Database {
	db.Delete(taskEntry, "")
	return db
}
