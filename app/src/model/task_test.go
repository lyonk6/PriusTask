package model

import (
	"fmt"
	"testing"
)

func TestGetTaskList(t *testing.T) {
	fmt.Println("TestGetTaskList")
	//func getTaskList(tt TaskTouch)
	// Presently taskLists are just a list of tasks ordered by due date,
	// thus we only need to test that the result returns a list of sorted
	// tasks.
	fmt.Println("Check 1")
	setTestDatabase()

	fmt.Println("Check 2")
	defer catchError("Error in TestGetTaskList ")
	tt := TaskTouch{}
	tt.ID = 1

	fmt.Println("Check 3")
	tl, err := getTaskList(tt)

	fmt.Println("Check 11")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Check 12")
	for i, v := range tl {
		fmt.Println(i, ". ", v)
	}
}

func TestUpdateTask(t *testing.T) {
	fmt.Println("TestUpdateTask")
	defer catchError("Error in TestUpdateTask ")
	task := Task{}
	updateTask(task)
}

func TestCreateTask(t *testing.T) {
	fmt.Println("TestCreateTask")
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
