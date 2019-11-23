package model

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
func getTaskList(tt TaskTouch) ([20]Task, error) {
	rows, err := db.Query(`SELECT id, userid, memo, repeatintervalindays, tasklength, duedate, creationdate, creationlongitude, creationlatitude FROM task ORDER BY DueDate ASC LIMIT 20;`)

	var tasks [20]Task
	t := &Task{}
	i := 0

	if err != nil {
		return tasks, err
	}

	//fmt.Print("Here are the columns: ")
	//fmt.Println(rows.Columns())

	for rows.Next() {
		err = rows.Scan(&t.ID, &t.UserID, &t.Memo, &t.RepeatIntervalInDays, &t.TaskLength, &t.DueDate, &t.CreationDate, &t.CreationLongitude, &t.CreationLatitude)

		if err != nil {
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
	DueDate              = $4
  WHERE id = $5`,
		t.Memo, t.RepeatIntervalInDays, t.TaskLength, t.DueDate, t.ID)
	return err
}

//Add a task to the Database and return the new Databases ID.
func createTask(t *Task) error {
	err := db.QueryRow(`
    INSERT INTO
    task(CreationDate,   CreationLatitude,   CreationLongitude,   DueDate,   Memo,   RepeatIntervalInDays,   TaskLength,   UserId)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		t.CreationDate, t.CreationLatitude, t.CreationLongitude, t.DueDate, t.Memo, t.RepeatIntervalInDays, t.TaskLength, t.UserID).Scan(&t.ID)
	return err
}

func (t *Task) toString() string {
	return t.Memo
}
