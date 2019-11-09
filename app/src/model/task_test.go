package model

import (
	"fmt"
	"testing"
)

func TestGetTaskList(t *testing.T) {
	//func getTaskList(tt TaskTouch)
	// Presently taskLists are just a list of tasks ordered by due date,
	// thus we only need to test that the result returns a list of sorted
	// tasks.
	fmt.Println("Check 1")
	setTestDatabase()
	fmt.Println(db.Ping())
	fmt.Println("Check 2")
	defer catchError("Error in TestGetTaskList ")
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
	}
}

func TestUpdateTask(t *testing.T) {
	defer catchError("Error in TestUpdateTask ")
	task := Task{}
	updateTask(task)
}

func TestCreateTask(t *testing.T) {
	defer catchError("Error in TestCreateTask ")
	task := Task{}
	createTask(task)
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
