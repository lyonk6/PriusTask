package model

//TaskTouch is an instance of a user updating or interacting with a Task.
type TaskTouch struct {
    ID                int32   `json:"id"`
    UserID            int32   `json:"userId"`
    TaskID            int32   `json:"taskId"`
    TouchTimeStamp    int64   `json:"touchTimeStamp"`
    LocationTimeStamp int64   `json:"locationTimeStamp"`
    Longitude         float32 `json:"longitude"`
    Latitude          float32 `json:"latitude"`
    Accuracy          float32 `json:"accuracy"`
    NetworkType       string  `json:"networkType"`
    TouchType         string  `json:"touchType"`
}
