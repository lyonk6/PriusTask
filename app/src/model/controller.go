package model

import (
    "encoding/json"
    "fmt"
    "net/http"
)

/*RegisterRoutes registers api method calls. Valid method calls include:
 *PostTaskTouch, GetTasks, PutTask & PostTask.
 */
func RegisterRoutes() {
    http.HandleFunc("/PostTaskTouch", func(w http.ResponseWriter, r *http.Request) {
        dec := json.NewDecoder(r.Body)
        var taskTouch TaskTouch
        err := dec.Decode(&taskTouch)
        fmt.Print("PostTaskTouch: ", err, " time: ", taskTouch.LocationTimeStamp)
        fmt.Print(" longitude: ", taskTouch.Longitude, " latitude: ", taskTouch.Latitude)
    })

    http.HandleFunc("/GetTasks", func(w http.ResponseWriter, r *http.Request) {
        dec := json.NewDecoder(r.Body)
        var taskTouch TaskTouch
        err := dec.Decode(&taskTouch)
        fmt.Print("GetTasks: ", err, " TODO: recieve a TaskTouch here.")
    })

    http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {
        dec := json.NewDecoder(r.Body)
        var task Task
        err := dec.Decode(&task)
        fmt.Println("GetTasks: ", err, " Body: ", task.Memo)
    })

    http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {
        dec := json.NewDecoder(r.Body)
        var task Task
        err := dec.Decode(&task)
        fmt.Println("GetTasks: err:", err, " Body: ", task.Memo)
    })
}
