package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Welcome to PriusTask!")
	http.ListenAndServe(":3000", nil)
	//http.ListenAndServeTLS(":3000", "certs/cert.pem", "certs/key.pem", nil)
}
