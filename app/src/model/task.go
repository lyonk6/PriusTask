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
	rows, err := db.Query(`SELECT * FROM $1.task ORDER BY DueDate DESC LIMIT 20;`, dbName)

	if err != nil {
		panic(err)
	}

	t := &Task{}
	for rows.Next() {
		err := rows.Scan(&t)

		if err != nil {
			panic(err)
		} else {
			t.toString()
		}
	}
}

//TODO update a task in the database.
func updateTask(t Task) {
	_, err := db.Exec(`
	UPDATE $1.task
	SET Memo = $2,
	RepeatIntervalInDays = $3,
	TaskLength           = $4,
	DueDate              = $5
    `)
	if err != nil {
		panic(err)
	}
	fmt.Println("updateTask: ", t.toString())
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
