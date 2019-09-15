package model


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
