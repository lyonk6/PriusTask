package model

import (
    "net/http/httptest"
    "net/http"
    "testing"
    "strings"
)

func TestRegisterRoutes(t *testing.T) {
    setTestDatabase()
    RegisterRoutes()
    //http.ListenAndServeTLS(":"+portNumber, "certs/cert.pem", "certs/key.pem", nil)
}

func dummyTaskTouch() TaskTouch {
  taskTouch := TaskTouch{}
  taskTouch.TouchType         = "START_UP"
  taskTouch.Accuracy          = 1.0
  taskTouch.Latitude          = 2.0
  taskTouch.Longitude         = 3.0
  taskTouch.LocationTimeStamp = 4
  taskTouch.ID                = 5
  taskTouch.TaskID            = 6
  taskTouch.TouchTimeStamp    = 7
  taskTouch.UserID            = 8
  return taskTouch
}

func dummyTask() Task {
  task := Task{}
  task.ID                   =-1
  task.UserID               = 0
  task.Memo                 = "Hello. "
  task.RepeatIntervalInDays = 1
  task.TaskLength           = 2
  task.DueDate              = 3
  task.CreationDate         = 4
  task.CreationLongitude    = 5
  task.CreationLatitude     = 6
  task.LastTouchType = "UPDATED"
  return task
}

func dummyEncodedTask() string {
  return `{"id":-1,"UserID":0,"memo":"Hello. ","repeatIntervalInDays":1,"taskLength":2,"dueDate":3,"creationDate":4,"creationLongitude":5,"creationLatitude":6,"lastTouchType":"UPDATED"}`
}

// Not a task. Has extra member field "derp":"flerp"
func dummyTotallyNotATask() string {
  return `{"id":-1,"UserID":0,"memo":"Hello. ","repeatIntervalInDays":1,"taskLength":2,"dueDate":3,"creationDate":4,"creationLongitude":5,"creationLatitude":6,"lastTouchType":"UPDATED","derp":"flerp"}`
}

// Not a taskTouch object.
func dummyTotallyNotATaskTouch() string {
  return `{ creationDate: 0, creationLatitude: 0, creationLongitude: 0, dueDate: 1618326000000, id: 112, lastTouchType: "", memo: "New Task Test 1", repeatIntervalInDays: 0, taskLength: 0, touchType: "DISMISSED", UserID: -1 }`
}

// A taskTouch object.
func dummyEncodedTaskTouch() string {
  return `{ "accuracy": 0.0, "id": 0, "latitude": 0.0, "locationTimeStamp": 0, "longitude": 0.0, "TaskID": -1, "touchTimeStamp": 1600729082413, "touchType": "START_UP", "UserID": 0 }`
}

func TestEncodeTask(t *testing.T) {
  //Step 1. Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()

  //Step 2. Create a task t:
  task := dummyTask()
  encodeTask(rr, &task)

  //Step 3. Validate the task was written to the Writer:\
  encoded  := strings.TrimSpace(rr.Body.String())
  expected := dummyEncodedTask()
  if encoded != expected {
    t.Errorf("handler returned unexpected body: got:\n %v \nwant:\n %v", encoded, expected)
    t.Errorf("got type: %T. Expect type: %T", encoded, expected)
  }
}

func TestDecodeTask(t *testing.T) {
  // (r *http.Request) Task
  // strings.NewReader(data.Encode())
  request, err := http.NewRequest("POST", "/PostTask", strings.NewReader(dummyEncodedTask()))
  if err != nil {
    t.Fatal(err)
  }

  task := decodeTask(request)
  if task != dummyTask(){
    t.Fatalf("Tasks are not equal: %v : %v", task, dummyTask())
  }

  // Mutate task and confirm that they no longer match:
  task.LastTouchType = "Invalid touch type :/"
  if task == dummyTask(){
    t.Fatalf("Tasks are equal and should not be: %v : %v", task, dummyTask())
  }

  // Confirm a task with an invalid touch type throws an error

  // Confirm that a task with inappropriate fields throws an exeception

}

func TestDecodeTaskTouchTouch(t *testing.T) {
  request, err := http.NewRequest("POST", "/PostTaskTouch", strings.NewReader(dummyEncodedTaskTouch()))
  if err != nil {
    t.Fatal(err)
  }

  taskTouch := decodeTaskTouch(request)
  if taskTouch != dummyTaskTouch(){
    t.Fatalf("TaskTouch are not equal: %v : %v", taskTouch, dummyTaskTouch())
  }

  // Mutate taskTouch and confirm that they no longer match:
  taskTouch.TouchType = "UPDATED"
  if taskTouch == dummyTaskTouch(){
    t.Fatalf("TaskTouch are equal and should not be: %v : %v", taskTouch, dummyTaskTouch())
  }

  // Confirm a taskTouch with an invalid touch type throws an error

  // Confirm that a taskTouch with inappropriate fields throws an exeception

}
