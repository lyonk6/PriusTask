package model

import (
	"fmt"
	"strconv"
)

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
	LastTouchType        string `json:"lastTouchType"`
}

//Return a list of tasks ordered by due date.
func getTaskList(tt TaskTouch) ([]Task, error) {
	rows, err := db.Query(`SELECT id, userid, memo, repeatintervalindays, tasklength, duedate, creationdate, creationlongitude, creationlatitude, lasttouchtype FROM task WHERE lasttouchtype <> $1 AND lasttouchtype <> $2 ORDER BY DueDate ASC LIMIT 10;`, "COMPLETED", "DELETED")

	//var tasks [20]Task
	tasks := make([]Task, 10)
	t := &Task{}
	i := 0

	if err != nil {
		fmt.Println("There was an error in the getTaskList query!!")
		return tasks, err
	}

	//fmt.Print("Here are the columns: ")
	//fmt.Println(rows.Columns())
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.UserID, &t.Memo, &t.RepeatIntervalInDays, &t.TaskLength, &t.DueDate, &t.CreationDate, &t.CreationLongitude, &t.CreationLatitude, &t.LastTouchType)

		if err != nil {
			fmt.Println("There was an error. Task could not be parsed from query.")
			return tasks, err
		}

		tasks[i] = *t
		i++
	}
	return tasks, err
}

//Update a task in the database.
func updateTask(t *Task) error {
	//fmt.Println("updateTask: ", t.toString())
	_, err := db.Exec(`
	UPDATE task
	SET memo             = $1,
	RepeatIntervalInDays = $2,
	TaskLength           = $3,
	DueDate              = $4,
  LastTouchType        = $5
  WHERE id = $6`,
		t.Memo, t.RepeatIntervalInDays, t.TaskLength, t.DueDate, t.LastTouchType, t.ID)
	return err
}

//Add a task to the Database and return the new Databases ID.
func createTask(t *Task) error {
	err := db.QueryRow(`
    INSERT INTO
    task(CreationDate,   CreationLatitude,   CreationLongitude,   DueDate,   Memo,   RepeatIntervalInDays,   TaskLength, LastTouchType, UserId)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		t.CreationDate, t.CreationLatitude, t.CreationLongitude, t.DueDate, t.Memo, t.RepeatIntervalInDays, t.TaskLength, t.LastTouchType, t.UserID).Scan(&t.ID)
	return err
}

func (t *Task) toString() string {
	return "ID: " + strconv.Itoa(int(t.ID)) + ", LastTouchType: " + t.LastTouchType +
		", Memo:" + t.Memo + ", RepeatInterval: " + strconv.FormatInt(t.RepeatIntervalInDays, 10) +
		", TaskLength: " + strconv.FormatInt(t.TaskLength, 10)
}
