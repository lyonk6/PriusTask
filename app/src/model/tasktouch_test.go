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
	//fmt.Println("Start by creating a task touch object. Then call saveTaskTouch()")
	tasktouch := &TaskTouch{}
	tasktouch.TouchType = "CREATED"
	tasktouch.Latitude = rand.Float64() * 31
	tasktouch.Longitude = rand.Float64() * 29

	err := saveTaskTouch(tasktouch)
	checkError(err)
	//fmt.Println("TaskTouch saved. Returned ID=" + strconv.FormatInt(int64(tasktouch.ID), 10))

	// Then delete the tasktouch.
	//fmt.Println("Now clean up the database by removing the recently created TaskTouch. ")
	statement := `DELETE FROM tasktouch WHERE id='` + strconv.FormatInt(int64(tasktouch.ID), 10) + `'`
	_, err = db.Exec(statement)
	checkError(err) //*/
}

func TestTouchTask(t *testing.T) {
	// func touchTask(tt *TaskTouch) error
	//First pick 5 tasks and set them to CREATED
	a := [5]int32{1, 2, 3, 4, 5}

	for i, v := range a {
		tt := TaskTouch{}
		tt.TaskID = v
		tt.TouchType = "CREATED"
		//TODO Loop though the local test database and mark all of these
		// tasks as deleted using the touchTask method.
		touchTask(&tt)
	}

	// TODO Refactor tests from TestPostTaskTouch here.

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
	//First create a dummy task and put it in the database. Give it a repeat interval.
	task.Memo = "This is a 'test' task!"
	task.LastTouchType = "CREATED"
	task.RepeatIntervalInDays = 7
	err := createTask(&task)
	checkError(err)

	startingTime := task.DueDate
	deltaTime := task.RepeatIntervalInDays * 24 * 60 * 60 * 1000

	//Create an appropriate task touch that does not modify the task.
	// Then call PostTaskTouch to post it.
	tasktouch.TaskID = task.ID
	tasktouch.TouchType = task.LastTouchType // "CREATED"
	err = postTaskTouch(&tasktouch)
	checkError(err)

	// Update the task type and optionally print what we have so far:
	//fmt.Println("Here is the original task and the new TouchType: ")
	tasktouch.TouchType = "COMPLETED"
	//fmt.Println("\tThe Task ID is:             ", task.ID)
	//fmt.Println("\tThe DueDate is:             ", task.DueDate)
	//fmt.Println("\tThe RepeatInterval is:      ", task.RepeatIntervalInDays)
	//fmt.Println("\tThe Task LastTouchType is:  ", task.LastTouchType)
	//fmt.Println("\tThe TaskTouch TouchType is: ", tasktouch.TouchType) //*/

	// Now Mark the Task and it's partner TaskTouch as completed. The API should
	// assign a new due date to this task.
	err = postTaskTouch(&tasktouch)
	checkError(err)

	//fmt.Println("Here is the updated task: ")
	// The DueDate on the task should now be 6.048(10)^8v (1 week in milliseconds.)
	stmt := `SELECT repeatintervalindays, duedate, lasttouchtype FROM task WHERE id=$1`
	err = db.QueryRow(stmt, task.ID).
		Scan(&task.RepeatIntervalInDays, &task.DueDate, &task.LastTouchType)
	checkError(err)

	if startingTime+deltaTime != task.DueDate {
		fmt.Println("Task ID: ", task.ID)
		fmt.Println("Expected Starting updated DueDate: ", startingTime+deltaTime)
		fmt.Println("Actual DueDate: ", task.DueDate)
		panic("Opps. These don't match")
	}

	//fmt.Println("The DueDate is:             ", task.DueDate)
	//fmt.Println("Repeat interval is:         ", task.RepeatIntervalInDays)
	//fmt.Println("The Task LastTouchType is:  ", task.LastTouchType)
	//fmt.Println("The TaskTouch TouchType is: ", tasktouch.TouchType) //*/

	// clean up this test by removing the task in question.
	_, err = db.Exec(`DELETE FROM task WHERE id=$1`, task.ID)
	checkError(err)
	//_, err = db.Exec(`UPDATE task SET lasttouchtype=$1 WHERE id=$2`, tt.TouchType, task.ID)

}
