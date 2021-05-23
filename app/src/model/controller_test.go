package model

import (
    "net/http/httptest"
    "testing"
    "strings"
)

func TestRegisterRoutes(t *testing.T) {
    setTestDatabase()
    RegisterRoutes()
    //http.ListenAndServeTLS(":"+portNumber, "certs/cert.pem", "certs/key.pem", nil)
}

/* An example test function:
func TestExampleFunc(t *testing.T) {
    r,_ := http.NewRequest("GET", "hello/1", nil)
    w := httptest.NewRecorder()
   //create a map of variable and set it into mux
    vars := map[string]string{
    "parameter_name": "parametervalue",
    }

   r = mux.SetURLVars(r, vars)
  callfun(w,r)
}*/

func TestEncodeTask(t *testing.T) {
    //(w http.ResponseWriter, t *Task)
  // Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()

    //Step 2. Create a task t:
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

  encodeTask(rr, &task)

    //Step 3. Validate the task was written to the Writer:\
  encoded  := strings.TrimSpace(rr.Body.String())
  expected := `{"id":-1,"userId":0,"memo":"Hello. ","repeatIntervalInDays":1,"taskLength":2,"dueDate":3,"creationDate":4,"creationLongitude":5,"creationLatitude":6,"lastTouchType":"UPDATED"}`
  if encoded != expected {
    t.Errorf("handler returned unexpected body: got:\n %v want:\n %v", rr.Body.String(), expected)
  }
}

func TestEncodeTaskList(t *testing.T) {
    //(w http.ResponseWriter, tl *[]Task)
}

func TestDecodeTask(t *testing.T) {
    //(r *http.Request) Task
}

func TestDecodeTaskTouch(t *testing.T) {
    //(r *http.Request) TaskTouch
}
