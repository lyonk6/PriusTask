package model

import (
	"fmt"
	"strconv"
)

var TouchTypes = map[string]bool{
	"UPDATED":    true,
	"DELETED":    true,
	"COMPLETED":  true,
	"DISMISSED":  true,
	"START_UP":   true,
	"HEART_BEAT": true,
	"CREATED":    true,
}

//TaskTouch is an instance of a user updating or interacting with a Task.
type TaskTouch struct {
	ID                int32   `json:"id"`
	UserID            int32   `json:"userId"`
	TaskID            int32   `json:"taskId"`
	TouchTimeStamp    int64   `json:"touchTimeStamp"`
	LocationTimeStamp int64   `json:"locationTimeStamp"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	Accuracy          float64 `json:"accuracy"`
	NetworkType       string  `json:"networkType"`
	TouchType         string  `json:"touchType"`
}

/**
 * Save a tasktouch to the database. Use the returned ID from the database
 * to give this tasktouch an ID.
 */
func saveTaskTouch(tt *TaskTouch) error {
	// TODO implement validation that the tt has a valid UserID and TaskID.
	// TODO also save an instance of this task with the task touch.

	err := db.QueryRow(`
        INSERT INTO tasktouch(UserID, TaskID, TouchTimeStamp, LocationTimeStamp, Longitude, Latitude, Accuracy, NetworkType, TouchType)
	    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		tt.UserID, tt.TaskID, tt.TouchTimeStamp, tt.LocationTimeStamp, tt.Longitude, tt.Latitude, tt.Accuracy, tt.NetworkType, tt.TouchType).Scan(&tt.ID)
	return err
}

/**
 * Save a new TaskTouch object and update a task according to how it was
 * modified. If a task is marked "COMPLETED" check and see if the task
 * repeats. If the task does repeat, mark the task as "UPDATED" and update
 * the new DueDate.
 *
 * All other updates should just modify the task.LastTouchType field.
 * Deleted tasks should not be removed yet.
 */
func postTaskTouch(tt *TaskTouch) error {
	var err error
	err = saveTaskTouch(tt)
	if err != nil {
		fmt.Println("Error returned when saving task touch.")
		return err
	}
	/* PostTaskTouch does not make a query if the touch type is
	 * "START_UP", "HEART_BEAT", or "CREATED".
	 */
	if tt.TouchType == "COMPLETED" || tt.TouchType == "DISMISSED" ||
		tt.TouchType == "DELETED" || tt.TouchType == "UPDATED" {
		err = touchTask(tt)
	}
	return err
}

/**
* Update task that has been modified as "UPDATED", "DELETED", "COMPLETED",
* or  "DISMISSED". Mark the LastTouchType accordingly and update the task
* due date if it is a repeating task.
*
 */
func touchTask(tt *TaskTouch) error {
	var err error
	var task Task
	err = db.QueryRow(`SELECT id, duedate, repeatintervalindays, memo, lasttouchtype FROM task WHERE id=$1`, tt.TaskID).
		Scan(&task.ID, &task.DueDate, &task.RepeatIntervalInDays, &task.Memo, &task.LastTouchType)

	if err != nil {
		return err
	}

	// Next see if we have a completed task that repeats. If so, set a
	// new dueDate and save the task as "UPDATED"
	if tt.TouchType == "COMPLETED" && task.RepeatIntervalInDays > 0 {

		// It is a repeating task and was just completed.
		task.DueDate = task.DueDate + (task.RepeatIntervalInDays * 86400000)
		_, err = db.Exec(`UPDATE task SET duedate=$1, lasttouchtype=$2 WHERE id=$3`,
			task.DueDate, "UPDATED", task.ID)

		// Otherwise, set the new update type.
		if err != nil {
			return err
		}
	} else {

		// Check if it is a reapeating task but was not completed.
		_, err = db.Exec(`UPDATE task SET lasttouchtype=$1 WHERE id=$2`, tt.TouchType, task.ID)
	}
	// Finally, return our error if we have one.
	return err
}

func (tt *TaskTouch) toString() string {
	time := strconv.FormatInt(tt.TouchTimeStamp, 10)
	long := strconv.FormatFloat(tt.Longitude, 'f', -1, 64)
	lati := strconv.FormatFloat(tt.Latitude, 'f', -1, 64)

	return "Date: " + time + " Longitude: " + long + " Latitude: " + lati + " Type: " + tt.TouchType
}
