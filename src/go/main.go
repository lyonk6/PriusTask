package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

func main() {
	portNumber, _ := getParameters()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("\"Welcome to PriusTask!\""))
		//w.Write([]byte(r.Body))
		//fmt.Printf("%+v\n", r)

		// DumpRequest for debugging purposes.
		DumpRequest(w, r)
	})

	http.ListenAndServeTLS(":"+portNumber, "certs/cert.pem", "certs/key.pem", nil)
}

/**
 * DumpRequest is used to dump an incomming request to CLI. This is helpful
 * for debugging REST calls. *
 */
func DumpRequest(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		//fmt.Fprint(w, err.Error())
		fmt.Println(string(requestDump))
	} else {
		//fmt.Fprint(w, string(requestDump))
		fmt.Println(string(requestDump))
	}
}

/**
 * Get the desired port numer and database url from a private paramerters
 * file. Return the port number as an int and the databaseUrl as a string
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
