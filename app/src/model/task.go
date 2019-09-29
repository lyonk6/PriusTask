package model

import "fmt"

// Task is an Object for holding a task.
type Task struct {
	ID                   int32  `json:"id"`
	UserID               int32  `json:"userId"`
	Memo                 string `json:"memo"`
	RepeatIntervalInDays int64  `json:"repeatIntervalInDays"`
	TaskLength           int64  `json:"taskLength"`
	DueDate              int64  `json:"dueDate"`
	CreationDate         int64  `json:"creationDate"`
	CreationLongitude    int64  `json:"creationLongitude"`
	CreationLatitude     int64  `json:"creationLatitude"`
}

//TODO return a list of tasks ordered by due date.
func getTaskList(tt TaskTouch) {
	fmt.Println("getTaskList: ", tt.toString())
}

//TODO update a task in the database.
func updateTask(t Task) {
	fmt.Println("updateTask: ", t.toString())
}

//TODO add a task to the Database.
func createTask(t Task) {
	fmt.Println("createTask: ", t.toString())
}

func (t *Task) toString() string {
	return t.Memo
}
