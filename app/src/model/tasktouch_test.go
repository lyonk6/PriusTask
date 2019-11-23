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
 *
 *
 */
func TestPostTaskTouch(t *testing.T) {
	fmt.Println("TestPostTaskTouch")

}
