package model

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq" //driver for postgres
)

func setTestDatabase() {
	if db == nil {
		url := os.Getenv("PT_TEST_DB_URL")

		connectToTestDatabase(url)
	}
}

func connectToTestDatabase(url string) {
	database, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	SetDatabase(database)
}

// Test the toString method for tasks and tasktouch objects.
func TestToString(t *testing.T) {
	task := Task{}
	task.Memo = "Here is a task."
	fmt.Println(task.toString())

	tasktouch := TaskTouch{}
	tasktouch.TouchType = "CREATED"
	fmt.Println(tasktouch.toString())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
