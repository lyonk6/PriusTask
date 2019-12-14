package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func encodeTask(w http.ResponseWriter, t *Task) {
	enc := json.NewEncoder(w)
	err := enc.Encode(t)
	if err != nil {
		panic(err)
	} else {
		return
	}
}

func encodeTaskList(w http.ResponseWriter, tl *[]Task) {
	enc := json.NewEncoder(w)
	err := enc.Encode(tl)
	if err != nil {
		panic(err)
	} else {
		return
	}
}

func decodeTask(r *http.Request) Task {
	dec := json.NewDecoder(r.Body)
	var t Task
	err := dec.Decode(&t)
	if err != nil {
		panic(err)
	} else {
		return t
	}
}

func decodeTaskTouch(r *http.Request) TaskTouch {
	dec := json.NewDecoder(r.Body)
	var tt TaskTouch
	err := dec.Decode(&tt)
	if err != nil {
		panic(err)
	} else {
		return tt
	}
}

/*RegisterRoutes registers api method calls. Valid method calls include:
 *PostTaskTouch, GetTasks, PutTask & PostTask.
 */
func RegisterRoutes() {
	http.HandleFunc("/PostTaskTouch", func(w http.ResponseWriter, r *http.Request) {
		tt := decodeTaskTouch(r)
		fmt.Print("PostTaskTouch- time:", tt.toString())
		postTaskTouch(&tt)
		getTaskList(tt)
	})

	// Call updateTask in task.go to update a task in the database.
	http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {
		t := decodeTask(r)
		fmt.Println("PutTask- Body: ", t.toString())
		err := updateTask(&t)
		clearError(err)
	})

	//Call creatTask in task.go to add a task to the database.
	http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {
		t := decodeTask(r)
		fmt.Println("PostTask- Body: ", t.toString())
		err := createTask(&t)
		clearError(err)
	})

	//All other requests get dumped.
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

func clearError(err error) {
	fmt.Println("\nAn error has been encountered:")
	fmt.Print(err)
}
