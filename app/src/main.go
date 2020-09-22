package main

import (
	"database/sql"
	"fmt"
	"model"
	"net/http"
	"os"

	_ "github.com/lib/pq" //driver for postgres
)

/**
 * Call getParameters() to retrieve the portNumber and database url. Then,
 * use connectToDatabase() to establish a connection and pass that dependency
 * the model. Then register our http request routes anf finally listen and
 * serve these requests.
 *
 *   func ListenAndServe(addr string, handler Handler) error
 *   func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
 */
func main() {
	portNumber := os.Getenv("PORT_NUMBER")
	url := os.Getenv("DB_URL")
	//portNumber, url := getParameters()
	connectToDatabase(url)
	model.RegisterRoutes()
	fmt.Println("Listening on port ", portNumber)
	fmt.Println("database URL: ", url)
	// Run the api insecurely:
	http.ListenAndServe(":"+portNumber, nil)

	// Run with TLS certificates like responsible adults:
	//http.ListenAndServeTLS(":"+portNumber, "certs/cert.pem", "certs/key.pem", nil)
}

/**
 * Connect to a database and set it for the model to use.
 */
func connectToDatabase(url string) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	model.SetDatabase(db)
}
