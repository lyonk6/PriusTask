package model

import (
	"fmt"
	"strconv"
)

var touchTypes = []string{"SAVED", "DELETED", "COMPLETED", "DISMISSED", "START_UP", "HEART_BEAT", "CREATED"}

//TaskTouch is an instance of a user updating or interacting with a Task.
type TaskTouch struct {
	ID                int64   `json:"id"`
	UserID            int64   `json:"userId"`
	TaskID            int64   `json:"taskId"`
	TouchTimeStamp    int64   `json:"touchTimeStamp"`
	LocationTimeStamp int64   `json:"locationTimeStamp"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	Accuracy          float64 `json:"accuracy"`
	NetworkType       string  `json:"networkType"`
	TouchType         string  `json:"touchType"`
}

//TODO Save this task touch and an instance of a task to an S3 bucket or some other DB.
func saveTaskTouch(tt TaskTouch) {
	fmt.Println("saveTaskTouch: ", tt.toString())
}

func postTaskTouch(tt TaskTouch) {
	saveTaskTouch(tt)
	switch tt.TouchType {
	case "DELETED":
		//TODO Delete a task from the Database.
	case "COMPLETED":
		//TODO Mark a task as completed --
		//  Update the due date (due date + repeatIntervalInDays)
		//  Delete the task if it is not meant to repeat.
	}
	fmt.Println("postTaskTouch: ", tt.toString(), touchTypes[0])
}

func (tt *TaskTouch) toString() string {
	time := strconv.FormatInt(tt.TouchTimeStamp, 10)
	long := strconv.FormatFloat(tt.Longitude, 'f', -1, 64)
	lati := strconv.FormatFloat(tt.Latitude, 'f', -1, 64)
	return "Date: " + time + " Longitude: " + long + " Latitude: " + lati
}
