package model

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/lib/pq" //driver for postgres
)

func setTestDatabase() {
	if db == nil {
		url := getTestURL()
		connectToDatabase(url)
	}
}

func connectToDatabase(url string) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	SetDatabase(db)
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
		fmt.Println("database URL: ", mySlice[3])
		databaseURL = mySlice[2][7:]
	} else {
		panic("Malformed params file. database identifier 'testdb=' and database url expected.")
	}

	return databaseURL
}
