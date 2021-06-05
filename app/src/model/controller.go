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

func decodeTask(r *http.Request) Task {
	dec := json.NewDecoder(r.Body)
  dec.DisallowUnknownFields() // Force errors
	var t Task
	err := dec.Decode(&t)
	if err != nil {
		panic(err)
	} else {
		return t
	}
}

func decodeTaskTouch(r *http.Request) (TaskTouch, error){
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
		//fmt.Println("\nRequest: PostTaskTouch: ", r)
		dumpRequest(w, r)
		tt, err := decodeTaskTouch(r)

		var tl []Task
		err = postTaskTouch(&tt)
		printError(err)
		tl, err = getTaskList(tt) // Here is an error. :(
		printError(err)

		/* TODO Only return a TaskList sometimes.
		if tt.TouchType == "COMPLETED" || tt.TouchType == "DISMISSED" || tt.TouchType == "START_UP" || tt.TouchType == "HEART_BEAT" {
			fmt.Println("Only return a TL sometimes.")
		}//*/

		encodeTaskList(w, tl)
		fmt.Print("Done. Request was: PostTaskTouch- time:", tt.toString(), "\n\n")
		//w.Write([]byte("[]")) //*/
	})

	// Call updateTask in task.go to update a task in the database.
	http.HandleFunc("/PutTask", func(w http.ResponseWriter, r *http.Request) {
		t := decodeTask(r)
		fmt.Println("PutTask- Body: ", t.toString())
		err := updateTask(&t)
		printError(err)
		encodeTask(w, &t)

	}) //*/

	//Call creatTask in task.go to add a task to the database.
	http.HandleFunc("/PostTask", func(w http.ResponseWriter, r *http.Request) {
		t := decodeTask(r)
		fmt.Println("\nRequest: PostTask- Body: ", t.toString())
		err := createTask(&t)
		printError(err)
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
		fmt.Fprint(w, err.Error())
		fmt.Println(string(requestDump))
	} else {
		fmt.Fprint(w, string(requestDump))
		fmt.Println(string(requestDump))
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
