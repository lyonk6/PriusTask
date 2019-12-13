package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

/**
 * Test saveTaskTouch. Start by creating a tasktouch object and submitting
 * it to the database. Then Clean up by removing it from the database.
 *
 * // TODO implement validation that the tt has a valid TouchType, UserID and TaskID.
 * // TODO also save an instance of this task with the task touch.
 */
func TestSaveTaskTouch(t *testing.T) {
	fmt.Println("Start by creating a task touch object. Then call saveTaskTouch()")
	tasktouch := &TaskTouch{}
	tasktouch.TouchType = "CREATED"
	tasktouch.Latitude = rand.Float64() * 31
	tasktouch.Longitude = rand.Float64() * 29

	err := saveTaskTouch(tasktouch)
	checkError(err)
	fmt.Println("TaskTouch saved. Returned ID=" + strconv.FormatInt(tasktouch.ID, 10))

	// Then delete the tasktouch.
	fmt.Println("Now clean up the database by removing the recently created TaskTouch. ")
	statement := `DELETE FROM tasktouch WHERE id='` + strconv.FormatInt(tasktouch.ID, 10) + `'`
	_, err = db.Exec(statement)
	checkError(err) //*/
}

/**
 * 1. Create a Task and put it in the Task database. Set it's LastTouchType to
 * "CREATED".
 *
 * 2. Make a TaskTouch that references this Task. Use the same touchtype. Then
 * use postTaskTouch() to save the tasktouch. Validate that the tasktouch was
 * saved and that the task was updated.
 *
 * 3. Give the task a Repeat interval. Then set the TouchType of the TaskTouch
 * to "COMPLETED". Validate that the task is updated with a new duedate.
 *
 *
 *
 */
func TestPostTaskTouch(t *testing.T) {
	fmt.Println("TestPostTaskTouch")
	task := Task{}
	tasktouch := TaskTouch{}
	//Create a task and put it in the database:
	task.Memo = "This is a 'test' task!"
	task.LastTouchType = "CREATED"
	task.RepeatIntervalInDays = 7
	createTask(&task)

	//Create a TaskTouch for the task above:
	tasktouch.TaskID = task.ID
	tasktouch.TouchType = task.LastTouchType
	err := postTaskTouch(&tasktouch)
	checkError(err)
	fmt.Println("Here is the original task: ")
	fmt.Println("The TaskID is:              ", task.DueDate)
	fmt.Println("The Task ID is:             ", task.DueDate)
	fmt.Println("The DueDate is:             ", task.DueDate)
	fmt.Println("The Task LastTouchType is:  ", task.LastTouchType)
	fmt.Println("The TaskTouch TouchType is: ", tasktouch.TouchType)

	//Now Mark the Task and it's partner TaskTouch as completed.

	tasktouch.TouchType = "COMPLETED"
	err = postTaskTouch(&tasktouch)
	checkError(err)

	fmt.Println("Here is the updated task: ")
	// The DueDate on the task should now be 6.048(10)^8
	stmt := `SELECT id, repeatintervalindays, duedate, lasttouchtype from task`
	err = db.QueryRow(stmt+` limit 1`).
		Scan(&task.ID, &task.RepeatIntervalInDays, &task.DueDate, &task.LastTouchType)
	checkError(err)
	fmt.Println("The DueDate is:             ", task.DueDate)
	fmt.Println("Repeat interval is:         ", task.RepeatIntervalInDays)
	fmt.Println("The Task LastTouchType is:  ", task.LastTouchType)
	fmt.Println("The TaskTouch TouchType is: ", tasktouch.TouchType)
}
