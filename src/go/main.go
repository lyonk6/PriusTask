package main

import (
     "fmt"
     "net/http"
		 "strings"
)

func main(){
     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
         w.Write([]byte("\"Welcome to PriusTask!\""))
         // w.Write([]byte(r.Body))
         //fmt.Printf("%+v\n", r)
         //fmt.Printf("Here is the body:%T ", r.Body,  r.Body, "\n")
		 formatRequest(r)
    })

     http.ListenAndServeTLS(":3000", "certs/cert.pem", "certs/key.pem", nil)
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
 // Create return string
 var request []string

 // Add the request string
 url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
 request = append(request, url) // Add the host
 request = append(request, fmt.Sprintf("Host: %v", r.Host)) // Loop through headers
 for name, headers := range r.Header {
   name = strings.ToLower(name)
   for _, h := range headers {
     request = append(request, fmt.Sprintf("%v: %v", name, h))
   }
 }

 // If this is a POST, add post data
 if r.Method == "POST" {
    r.ParseForm()
    request = append(request, "\n")
    request = append(request, r.Form.Encode())
 }   // Return the request as a string
  return strings.Join(request, "\n")
}
