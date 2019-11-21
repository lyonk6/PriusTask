package model

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetTaskList(t *testing.T) {
	// func getTaskList(tt TaskTouch)
	// Presently taskLists are just a list of tasks ordered by due date,
	// thus we only need to test that the result returns a list of sorted
	// tasks.
	fmt.Println("Check 1")
	setTestDatabase()
	fmt.Println("Check 2")
	tt := TaskTouch{}
	tt.ID = 1

	fmt.Println("Check 3")
	tl, err := getTaskList(tt)

	fmt.Println("No unhandled exceptions in getTaskList")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No handled exceptions in getTaskList")
	}

	for i, v := range tl {
		fmt.Println(i, ". ", v)
		//fmt.Println("Here is a due date: ", tl[i])
		if i > 0 && tl[i].DueDate < tl[i-1].DueDate && tl[i].ID != 0 {
			panic("Tasks not sorted by due date!")
		}
	}
}

func TestCreateTask(t *testing.T) {
	task := Task{}
	task.Memo = "This is the story of a girl"

	createTask(&task)
	_, err := db.Query(`DELETE FROM task WHERE id="` + strconv.FormatInt(task.ID, 10) + `"`)
	if err != nil {
		panic(err)
	}
}

func TestUpdateTask(t *testing.T) {
	// First get a task, then make a change to it. Then verify the
	// change was made.
}

func TestToString(t *testing.T) {
	defer catchError("Error in TestToString ")
	task := Task{}
	task.Memo = "Here is a task."
	fmt.Println(task.toString())
}

func catchError(message string) {
	if r := recover(); r != nil {
		fmt.Println("task_test.catchError: ", message)
		fmt.Println(r)
		//debug.PrintStack()
	}
}
