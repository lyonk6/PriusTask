package model

import (
	"fmt"
)

// Task is an Object for holding a task.
type Task struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"userId"`
	Memo                 string `json:"memo"`
	RepeatIntervalInDays int64  `json:"repeatIntervalInDays"`
	TaskLength           int64  `json:"taskLength"`
	DueDate              int64  `json:"dueDate"`
	CreationDate         int64  `json:"creationDate"`
	CreationLongitude    int64  `json:"creationLongitude"`
	CreationLatitude     int64  `json:"creationLatitude"`
}

//Return a list of tasks ordered by due date.
func getTaskList(tt TaskTouch) ([]Task, error) {
	fmt.Println("Check 4")
	rows, err := db.Query(`SELECT id, userid, memo, repeatintervalindays, tasklength, duedate, creationdate, creationlongitude, creationlatitude FROM task ORDER BY DueDate DESC LIMIT 20;`)
	fmt.Println("Check 5")

	if err != nil {
		return nil, err
	}

	fmt.Println("Check 6")
	var tasks []Task
	t := &Task{}
	i := 0

	fmt.Println("Check 7")
	//fmt.Println(rows.Columns())

	for rows.Next() {
		fmt.Println("Check 8")
		err := rows.Scan(&t.ID, &t.UserID, &t.Memo, &t.RepeatIntervalInDays, &t.TaskLength, &t.DueDate, &t.CreationDate, &t.CreationLongitude, &t.CreationLatitude)

		fmt.Println("Check 9")
		if err != nil {
			panic(err)
		} else {
			fmt.Println(t.toString())

		}
		fmt.Println("Check 10") //It breaks right here.
		tasks[i] = *t
		i++
	}
	return tasks, nil
}

//Update a task in the database.
func updateTask(t Task) {
	fmt.Println("updateTask: ", t.toString())
	_, err := db.Exec(`
	UPDATE task
	SET Memo = $1,
	RepeatIntervalInDays = $2,
	TaskLength           = $3,
	DueDate              = $4
    `)
	if err != nil {
		panic(err)
	}
}

//TODO add a task to the Database.
func createTask(t Task) {
	result, err := db.Exec(`
    INSERT INTO
    task(CreationDate,   CreationLatitude,   CreationLongitude,   DueDate,   Memo,   RepeatIntervalInDays,   TaskLength,   UserId)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) `,
		t.CreationDate, t.CreationLatitude, t.CreationLongitude, t.DueDate, t.Memo, t.RepeatIntervalInDays, t.TaskLength, t.UserID)
	if err != nil {
		panic(err)
	}
	t.ID, _ = result.LastInsertId()
	fmt.Println("createTask: ", t.toString())
}

func (t *Task) toString() string {
	return t.Memo
}
