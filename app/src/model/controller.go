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
} //*/

/*RegisterRoutes registers api method calls. Valid method calls include:
 *PostTaskTouch, GetTasks, PutTask & PostTask.
 */
func RegisterRoutes() {
	// TODO Decide when a TaskList really needs to be updated.
	http.HandleFunc("/PostTaskTouch", func(w http.ResponseWriter, r *http.Request) {
		tt := decodeTaskTouch(r)
		fmt.Println("\nRequest: PostTaskTouch- time:", tt.toString())
		/*
			var tl *[]Task
			err := postTaskTouch(&tt)
			checkError(err)
			*tl, err = getTaskList(tt)
			checkError(err)
			encodeTaskList(w, tl)
			//*/

		//w.Write([]byte("200 Success"))
		w.Write([]byte("[]")) //*/
	})

	// Call updateTask in task.go to update a task in the database.
	http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {
		t := decodeTask(r)
		fmt.Println("PutTask- Body: ", t.toString())
		err := updateTask(&t)
		checkError(err)
		w.Write([]byte("{}"))
	}) //*/

	//Call creatTask in task.go to add a task to the database.
	http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {
		t := decodeTask(r)
		fmt.Println("\nRequest: PostTask- Body: ", t.toString())
		err := createTask(&t)
		checkError(err)
		encodeTask(w, &t)
		fmt.Println("Here is the task we are returning: " + t.toString())
		//w.Write([]byte("200 Success"))
	}) //*/

	//All other requests get dumped.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\n400 Bad Request")
		w.Write([]byte("400 Bad Request"))
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

/**
 * Used by tests and prod to validate no errors are returned from a function call.
 */
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
