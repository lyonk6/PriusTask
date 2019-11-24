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
 * use postTaskTouch() save the tasktouch. Validate that the tasktouch was saved
 * and that the task was updated.
 *
 * 3. Give the task a Repeat interval. Then set the TouchType of the TaskTouch
 * to "COMPLETED". Validate that the task is updated with a new duedate.
 */
func TestPostTaskTouch(t *testing.T) {
	fmt.Println("TestPostTaskTouch")

}
