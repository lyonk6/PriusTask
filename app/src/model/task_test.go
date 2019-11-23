package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// Presently taskLists are just a list of tasks ordered by due date,
// thus we only need to test that the result returns a list of sorted
// tasks.
func TestGetTaskList(t *testing.T) {
	// func getTaskList(tt TaskTouch)
	setTestDatabase()
	tt := TaskTouch{}
	tt.ID = 1

	tl, err := getTaskList(tt)

	fmt.Println("No unhandled exceptions in getTaskList")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No handled exceptions in getTaskList")
	}

	for i := range tl {
		//fmt.Println(i, ". ", v)
		//fmt.Println("Here is a due date: ", tl[i])
		if i > 0 && tl[i].DueDate < tl[i-1].DueDate && tl[i].ID != 0 {
			panic("Tasks not sorted by due date!")
		}
	}
}

/**
 * Create a task reference and call createTask() to receive a
 * task id. Use this id to delete the task.
 */
func TestCreateTask(t *testing.T) {
	task := Task{}
	task.Memo = "This is the story of a girl"
	// First add a new task to the database and check the error.
	err := createTask(&task)
	if err != nil {
		panic(err)
	}
	// Then delete the task.
	statement := `DELETE FROM task WHERE id='` + strconv.FormatInt(task.ID, 10) + `'`
	_, err = db.Exec(statement)
	if err != nil {
		panic(err)
	}
}

/**
 * First fetch an arbitrary task. Then make changes to the task and
 * call updateTask to change it in the database. Then verify the tasks
 * has been updated.
 */
func TestUpdateTask(t *testing.T) {
	// First fetch an arbitrary task.
	//fmt.Println("Fetch an arbitrary task and call it 'task' ...")
	task := &Task{}
	var original Task
	stmt := `SELECT id, userid, memo, repeatintervalindays, tasklength, duedate, creationdate, creationlongitude, creationlatitude, lasttouchtype from task`
	err := db.QueryRow(stmt+` limit 1`).
		Scan(&task.ID, &task.UserID, &task.Memo,
			&task.RepeatIntervalInDays, &task.TaskLength,
			&task.DueDate, &task.CreationDate, &task.CreationLongitude,
			&task.CreationLatitude, &task.LastTouchType)
	checkError(err)

	//fmt.Println("Then make a copy of it and call it 'original'...")
	original = *task
	//fmt.Println(*task)
	//fmt.Println(original)

	// Then make changes to the task (these are the only mutable values):
	task.Memo = "I am not a theorem in PM"
	task.RepeatIntervalInDays = 29
	task.TaskLength = 3600001 // 1 hour in ms
	task.DueDate = rand.Int63()
	task.LastTouchType = "START_UP"
	// Now add it to a database ...
	//fmt.Println(`Mutate "task" then making an update to the db ...`)
	err = updateTask(task)
	checkError(err)
	//fmt.Println(*task)
	//fmt.Println(original)

	// Fetch the same task from the database and confirm it has been updated.
	//fmt.Println(`Fetch the task again. This time call it 'updated' ...`)
	updated := &Task{}
	err = db.QueryRow(stmt+` WHERE id=$1`, task.ID).
		Scan(&updated.ID, &updated.UserID, &updated.Memo,
			&updated.RepeatIntervalInDays, &updated.TaskLength,
			&updated.DueDate, &updated.CreationDate, &updated.CreationLongitude,
			&updated.CreationLatitude, &updated.LastTouchType)
	checkError(err)

	//fmt.Println(`Finally verify the 'updated' task is the same as 'task'`)
	if *task != *updated {
		panic("Error! Task not updated in TestUpdateTask!!!")
	}

	//fmt.Println(task)
	//fmt.Println(original)

	// Finally restore the original value:
	err = updateTask(&original)
	checkError(err)
}

// Test the toString method for tasks.
func TestToString(t *testing.T) {
	task := Task{}
	task.Memo = "Here is a task."
	fmt.Println(task.toString())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
