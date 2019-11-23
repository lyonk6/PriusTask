package model

import (
	"fmt"
	"strconv"
)

var touchTypes = []string{"UPDATED", "DELETED", "COMPLETED", "DISMISSED", "START_UP", "HEART_BEAT", "CREATED"}

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

// Save a tasktouch to the database and give this taskTouch an ID.
func saveTaskTouch(tt *TaskTouch) error {
	fmt.Println("saveTaskTouch: ", tt.toString())
	err := db.QueryRow(`
        INSERT INTO tasktouch(UserID, TaskID, TouchTimeStamp, LocationTimeStamp, Longitude, Latitude, Accuracy, NetworkType, TouchType)
	    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		tt.UserID, tt.TaskID, tt.TouchTimeStamp, tt.LocationTimeStamp, tt.Longitude, tt.Latitude, tt.Accuracy, tt.NetworkType, tt.TouchType).Scan(&tt.ID)
	return err
}

func postTaskTouch(tt *TaskTouch) error {
	var err error
	var task Task
	saveTaskTouch(tt)

	// First check and see if the touch type is "COMPLETED"
	if tt.TouchType == "COMPLETED" {
		// Fetch the coresponding task that has been completed:
		err = db.QueryRow(`SELECT id, duedate, repeatintervalindays FROM task where id=$1`, tt.UserID).
			Scan(&task.ID, &task.DueDate, &task.RepeatIntervalInDays)

			// If the reapeat interval is > 0 then set a new dueDate and save the task as "UPDATED"
		if task.RepeatIntervalInDays > 0 && err == nil {
			_, err = db.Exec(`UPDATE task SET DueDate='$1', LastTouchType ='$2' WHERE id='$3'`,
				task.DueDate+task.RepeatIntervalInDays*86400000, "UPDATED", tt.ID)
		}
		// Otherwise, set the new update type.
	} else if err == nil {
		_, err = db.Exec(`UPDATE task SET lasttouchtype='$1' WHERE id='$2'`, tt.TouchType, tt.ID)
	}
	fmt.Println("postTaskTouch: ", tt.toString(), touchTypes[0])
	return err
}

func (tt *TaskTouch) toString() string {
	time := strconv.FormatInt(tt.TouchTimeStamp, 10)
	long := strconv.FormatFloat(tt.Longitude, 'f', -1, 64)
	lati := strconv.FormatFloat(tt.Latitude, 'f', -1, 64)

	return "Date: " + time + " Longitude: " + long + " Latitude: " + lati + " Type: " + tt.TouchType
}
