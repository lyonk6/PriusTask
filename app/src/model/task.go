package model

import "net/http"

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

func (t Task) registerRoutes() {
    http.HandleFunc("/GetTasks", func(w http.ResponseWriter, r *http.Request) {

    })

    http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {

    })

    http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {

    })
}
