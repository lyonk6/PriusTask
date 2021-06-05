package model

import (
	"database/sql"
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
	task.RepeatIntervalInDays = 1
	if task.toString() != "ID: 0, LastTouchType: , Memo:Here is a task., RepeatInterval: 1, TaskLength: 0" {
		t.Errorf("task.toString() does not work ")
	}

	tasktouch := TaskTouch{}
	tasktouch.TouchType = "CREATED"
	tasktouch.Latitude = 1
	if tasktouch.toString() != "Date: 0 Longitude: 0 Latitude: 1 Type: CREATED" {
		t.Errorf("tasktouch.toString() does not work ")
	}
}
