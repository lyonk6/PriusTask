package controller

import (
    "fmt"
    "net/http"
)

// RegisterRoutes registers api method calls.
func RegisterRoutes() {
    http.HandleFunc("/PostTaskTouch", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print("PostTaskTouch: ", r.Body)
    })

    http.HandleFunc("/GetTasks", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print("GetTasks: ", r.Body)
    })

    http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print("PutTask: ", r.Body)
    })

    http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print("PostTask: ", r.Body)
    })
}
