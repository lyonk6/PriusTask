package model

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	setTestDatabase()
	RegisterRoutes()
}

func dummyTaskTouch() TaskTouch {
	taskTouch := TaskTouch{}
	taskTouch.TouchType = "START_UP"
	taskTouch.Accuracy = 1.0
	taskTouch.Latitude = 2.0
	taskTouch.Longitude = 3.0
	taskTouch.LocationTimeStamp = 4
	taskTouch.ID = 5
	taskTouch.TaskID = 6
	taskTouch.TouchTimeStamp = 7
	taskTouch.UserID = 8
	return taskTouch
}

func dummyTask() Task {
	task := Task{}
	task.ID = -1
	task.UserID = 0
	task.Memo = "Hello. "
	task.RepeatIntervalInDays = 1
	task.TaskLength = 2
	task.DueDate = 3
	task.CreationDate = 4
	task.CreationLongitude = 5
	task.CreationLatitude = 6
	task.LastTouchType = "UPDATED"
	return task
}

func dummyEncodedTask() string {
	return `{"id":-1,"userId":0,"memo":"Hello. ","repeatIntervalInDays":1,"taskLength":2,"dueDate":3,"creationDate":4,"creationLongitude":5,"creationLatitude":6,"lastTouchType":"UPDATED"}`
}

// Not a task. Has extra member field "derp":"flerp"
func dummyTotallyNotATask() string {
	return `{"id":-1,"userId":0,"memo":"Hello. ","repeatIntervalInDays":1,"taskLength":2,"dueDate":3,"creationDate":4,"creationLongitude":5,"creationLatitude":6,"lastTouchType":"UPDATED","derp":"flerp"}`
}

// Not a taskTouch object.
func dummyTotallyNotATaskTouch() string {
	return `{ ,, "latitude": 2.0, "locationTimeStamp": 4, "longitude": 3.0, "taskId": 6, "touchTimeStamp": 7, "touchType": "START_UP", "userId": 8 }`
}

// A taskTouch object.
func dummyEncodedTaskTouch() string {
	return `{ "accuracy": 1.0, "id": 5, "latitude": 2.0, "locationTimeStamp": 4, "longitude": 3.0, "taskId": 6, "touchTimeStamp": 7, "touchType": "START_UP", "userId": 8 }`
}

func TestEncodeTask(t *testing.T) {
	//Step 1. Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	//Step 2. Create a task t:
	task := dummyTask()
	encodeTask(rr, &task)

	//Step 3. Validate the task was written to the Writer:\
	encoded := strings.TrimSpace(rr.Body.String())
	encoded = strings.TrimSuffix(encoded, "\n")
	expected := dummyEncodedTask()
	if encoded != expected {
		t.Errorf("handler returned unexpected body: got:\n %v \nwant:\n %v", encoded, expected)
		t.Errorf("got type: %T. Expect type: %T", encoded, expected)
		t.Errorf("got type: %v. Expect type: %v", len(encoded), len(expected))
	}
}

func TestDecodeTask(t *testing.T) {
	request, err := http.NewRequest("POST", "/PostTask", strings.NewReader(dummyEncodedTask()))
	if err != nil {
		t.Fatal(err)
	}

	task := decodeTask(request)
	if task != dummyTask() {
		t.Fatalf("Tasks are not equal: %v : %v", task, dummyTask())
	}

	// Mutate task and confirm that they no longer match:
	task.LastTouchType = "Invalid touch type :/"
	if task == dummyTask() {
		t.Fatalf("Tasks are equal and should not be: %v : %v", task, dummyTask())
	}

	// Confirm a task with an invalid touch type throws an error

	// Confirm that a task with inappropriate fields throws an exeception

}

func TestDecodeTaskTouchTouch(t *testing.T) {
	// 1. Valid case:
	// Create a new POST TaskTouch request use "dummyEncodedTaskTouch" and check for errors.
	request, err := http.NewRequest("POST", "/PostTaskTouch", strings.NewReader(dummyEncodedTaskTouch()))
	if err != nil {
		t.Fatal(err)
	}

	// Next, decode the tasktouch object and validate it is the same as our dummy TaskTouch object:
	taskTouch, _ := decodeTaskTouch(request)
	if taskTouch != dummyTaskTouch() {
		t.Fatalf("TaskTouch are not equal: %v : %v", taskTouch, dummyTaskTouch())
	}

	// For sanity sake, mutate the taskTouch and confirm that they no longer match:
	taskTouch.TouchType = "UPDATED"
	if taskTouch == dummyTaskTouch() {
		t.Fatalf("TaskTouch are equal and should not be: %v : %v", taskTouch, dummyTaskTouch())
	}

	// 2. Error case: Confirm that a taskTouch with inappropriate fields throws an exeception.
	request, err = http.NewRequest("POST", "/PostTaskTouch", strings.NewReader(dummyEncodedTask()))
	if err != nil {
		t.Fatal(err)
	}

	taskTouch, err = decodeTaskTouch(request)
	if err == nil {
		t.Fatalf("Fatal Error. An invalid object should not be decoded. TaskTouch : %v", taskTouch.toString())
	}

	request, err = http.NewRequest("POST", "/PostTaskTouch", strings.NewReader(dummyTotallyNotATaskTouch()))
	if err != nil {
		t.Fatal(err)
	}

	taskTouch, err = decodeTaskTouch(request)
	if err == nil {
		t.Fatalf("Fatal Error. An invalid object should not be decoded. TaskTouch : %v", taskTouch.toString())
	}
}
