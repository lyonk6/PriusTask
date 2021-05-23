package model

import (
    "net/http/httptest"
    "testing"
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
	task.Memo = "Hello. "
  task.LastTouchType = "UPDATED"

  encodeTask(rr, &task)

	//Step 3. Validate the task was written to the Writer:

	//encodeTask(w, t)
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
