package database

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

type (
	Task struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Active      bool   `json:"active"`
		Time        int    `json:"time"`
		Username    string `json:"username"`
	}

	Database *scribble.Driver
)

const TaskEntry string = "task"
const dir string = "./database/.data"

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

func SaveTask(entry string, task Task) Database {
	return save(entry, "test-file", task)
}
