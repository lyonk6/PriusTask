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

//Slices have a pointer already so don't pass by reference.
func encodeTaskList(w http.ResponseWriter, tl []Task) {
	enc := json.NewEncoder(w)
	err := enc.Encode(tl)
	if err != nil {
		panic(err)
	} else {
		return
	}
}

func decodeTask(r *http.Request) (Task, error) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // Force errors
	var t Task
	err := dec.Decode(&t)
	return t, err
}

func decodeTaskTouch(r *http.Request) (TaskTouch, error) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // Force errors
	var tt TaskTouch
	err := dec.Decode(&tt)
	return tt, err
} //*/

/*RegisterRoutes registers api method calls. Valid method calls include:
 *PostTaskTouch, GetTasks, PutTask & PostTask.
 */
func RegisterRoutes() {
	http.HandleFunc("/PostTaskTouch", func(w http.ResponseWriter, r *http.Request) {
		// First try to decode the TaskTouch object. Check for an error.
		tt, err := decodeTaskTouch(r)
		if err != nil {
			respondBadRequest(w, r)
		} else {
			var tl []Task
			err = postTaskTouch(&tt)
			printError(err)
			tl, err = getTaskList(tt)
			printError(err)

			if tt.TouchType == "COMPLETED" || tt.TouchType == "DISMISSED" || tt.TouchType == "START_UP" || tt.TouchType == "HEART_BEAT" {
				encodeTaskList(w, tl)
			} //*/
		}
	})

	// Call updateTask in task.go to update a task in the database.
	http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {
		t, err := decodeTask(r)
		if err != nil {
			respondBadRequest(w, r)
		} else {
			err = updateTask(&t)
			printError(err)
			encodeTask(w, &t)
		}
	}) //*/

	//Call creatTask in task.go to add a task to the database.
	http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {
		t, err := decodeTask(r)
		if err != nil {
			respondBadRequest(w, r)
		} else {
			err = createTask(&t)
			printError(err)
			encodeTask(w, &t)
		}
	}) //*/

	// All other requests get dumped.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respondBadRequest(w, r)
	})
}

func respondBadRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n400 Bad Request")
	w.Write([]byte("400 Bad Request"))
	dumpRequest(r)
}

/**
 * DumpRequest is used to dump an incomming request to CLI. This is helpful
 * for debugging REST calls. *
 */
func dumpRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(string(requestDump))
	} else {
		fmt.Println(string(requestDump))
	}
}

/**
 * Check if there's an error. If so, panic.
 */
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * Used by tests and prod to validate no errors are returned from a function call.
 */
func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
