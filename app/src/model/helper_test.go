package model

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	_ "github.com/lib/pq" //driver for postgres
)

func setTestDatabase() {
	if db == nil {
		url := getTestURL()
		connectToTestDatabase(url)
	}
}

func connectToTestDatabase(url string) {
	database, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	SetDatabase(database)
}

func getTestURL() (databaseURL string) {
	file, err := ioutil.ReadFile("../../../params")

	//verify the file is not null:
	if err != nil {
		panic(err)
	}

	//Split the file into a slice of lines:
	mySlice := strings.Split(string(file), "\n")

	//Verify prefix for line 2 and assign databaseURL:
	if strings.HasPrefix(mySlice[2], "testdb=") && len(mySlice[2]) > 13 {
		//fmt.Println("test database URL: ", mySlice[3])
		databaseURL = mySlice[2][7:]
	} else {
		panic("Malformed params file. database identifier 'testdb=' and database url expected.")
	}
	return databaseURL
}

/**
 * Used by test files to validate no errors are returned from calls
 * made to the database.
 */
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Test the toString method for tasks and tasktouch objects.
func TestToString(t *testing.T) {
	task := Task{}
	task.Memo = "Here is a task."
	fmt.Println(task.toString())

	tasktouch := TaskTouch{}
	tasktouch.TouchType = "CREATED"
	fmt.Println(tasktouch.toString())
}
