package model

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httputil"
)

func decodeTask(r *http.Request) Task {
    dec := json.NewDecoder(r.Body)
    var task Task
    err := dec.Decode(&task)
    if err != nil {
        panic(err)
    } else {
        return task
    }
}

func decodeTaskTouch(r *http.Request) TaskTouch {
    dec := json.NewDecoder(r.Body)
    var taskTouch TaskTouch
    err := dec.Decode(&taskTouch)
    if err != nil {
        panic(err)
    } else {
        return taskTouch
    }
}

/*RegisterRoutes registers api method calls. Valid method calls include:
 *PostTaskTouch, GetTasks, PutTask & PostTask.
 */
func RegisterRoutes() {
    http.HandleFunc("/PostTaskTouch", func(w http.ResponseWriter, r *http.Request) {
        taskTouch := decodeTaskTouch(r)
        fmt.Print("PostTaskTouch- time:", taskTouch.LocationTimeStamp)
        fmt.Println(" longitude: ", taskTouch.Longitude, " latitude: ", taskTouch.Latitude)
    })

    http.HandleFunc("/GetTasks", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("GetTasks- TODO: recieve a TaskTouch here.")
    })

    http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {
        task := decodeTask(r)
        fmt.Println("GetTasks- Body: ", task.Memo)
    })

    http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {
        task := decodeTask(r)
        fmt.Println("GetTasks- Body: ", task.Memo)
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("400 Bad Request"))
        fmt.Println("400 Bad Request")
        dumpRequest(w, r)
    })

}

/**
 * DumpRequest is used to dump an incomming request to CLI. This is helpful
 * for debugging REST calls. *
 */
func dumpRequest(w http.ResponseWriter, r *http.Request) {
    requestDump, err := httputil.DumpRequest(r, true)
    if err != nil {
        //fmt.Fprint(w, err.Error())
        fmt.Println(string(requestDump))
    } else {
        //fmt.Fprint(w, string(requestDump))
        fmt.Println(string(requestDump))
    }
}

// SetDatabase sets the database for this package.
func SetDatabase(db *sql.DB) {
    fmt.Println("Database not set.")
}
