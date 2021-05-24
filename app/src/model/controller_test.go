package model

import (
    "net/http/httptest"
    "net/http"
    "testing"
    "strings"
    //"fmt"
)

func TestRegisterRoutes(t *testing.T) {
    setTestDatabase()
    RegisterRoutes()
    //http.ListenAndServeTLS(":"+portNumber, "certs/cert.pem", "certs/key.pem", nil)
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
  return `{"id":-1,"userId":0,"memo":"Hello. ","repeatIntervalInDays":1,"taskLength":2,"dueDate":3,"creationDate":4,"creationLongitude":5,"creationLatitude":6,"lastTouchType":"UPDATED"}`
}

// Not a task. Has extra member field "derp":"flerp"
func dummyTotallyNotATask() string {
  return `{"id":-1,"userId":0,"memo":"Hello. ","repeatIntervalInDays":1,"taskLength":2,"dueDate":3,"creationDate":4,"creationLongitude":5,"creationLatitude":6,"lastTouchType":"UPDATED","derp":"flerp"}`
}

func TestEncodeTask(t *testing.T) {
  //(w http.ResponseWriter, t *Task)
  // Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()

  //Step 2. Create a task t:
  task := dummyTask()
  encodeTask(rr, &task)

  //Step 3. Validate the task was written to the Writer:\
  encoded  := strings.TrimSpace(rr.Body.String())
  expected := dummyEncodedTask()
  if encoded != expected {
    t.Errorf("handler returned unexpected body: got:\n %v want:\n %v", rr.Body.String(), expected)
  }
}

func TestEncodeTaskList(t *testing.T) {
    //(w http.ResponseWriter, tl *[]Task)
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

  // Confirm that a tasks with inappropriate fields throws an exeception

}

func TestDecodeTaskTouch(t *testing.T) {
    //(r *http.Request) TaskTouch
}
