package model

import (
	"fmt"
	"testing"
)

func otherFunction() {
}

func TestGetTaskList(t *testing.T) {
	fmt.Println("TestGetTaskList")
	//func getTaskList(tt TaskTouch)

	// Presently taskLists are just a list of tasks ordered by due date,
	// thus we only need to test that the result returns a list of sorted
	// tasks.
	//tt := TaskTouch{}
	//tt.ID = 1
	//getTaskList(tt)
}

func TestUpdateTask(t *testing.T) {
	fmt.Println("TestUpdateTask")

}

func TestCreateTask(t *testing.T) {
	fmt.Println("TestCreateTask")
}
