package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq" //driver for postgres
)

func main() {
	portNumber, url := getParameters()
	connectToDatabase(url)
	model.RegisterRoutes()
	http.ListenAndServeTLS(":"+portNumber, "certs/cert.pem", "certs/key.pem", nil)
}

func connectToDatabase(url string) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println(err)
	}
	model.SetDatabase(db, "priustask")
}

/**
 * Get the desired port numer and database url from a private paramerters
 * file. Return the port number and the databaseUrl as a string.
 */
func getParameters() (portNumber, databaseURL string) {
	file, err := ioutil.ReadFile("params")

	//verify the file is not null:
	if err != nil {
		panic(err)
	}

	//Split the file into a slice of lines:
	mySlice := strings.Split(string(file), "\n")

	//Test if the port number specified is a properly formatted number:
	_, err = strconv.Atoi(mySlice[0][5:])
	if strings.HasPrefix(mySlice[0], "port=") && err == nil {
		//port_number = mySlice[1][5:]
		portNumber = mySlice[0][5:]
		fmt.Println("Listening on port ", portNumber)
	} else {
		panic("Malformed params file. port identifier 'port=' and port number expected.")
	}

	//Verify prefix for line 2 and assign databaseURL:
	if strings.HasPrefix(mySlice[1], "db=") && len(mySlice[1]) > 9 {
		fmt.Println("database URL: ", mySlice[1])
		databaseURL = mySlice[1][3:]
	} else {
		panic("Malformed params file. database identifier 'db=' and database url expected.")
	}

	return portNumber, databaseURL
}
