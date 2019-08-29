package main

import (
     "fmt"
     "net/http"
     "net/http/httputil"
)

func main(){
     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
       w.Write([]byte("\"Welcome to PriusTask!\""))
	    //w.Write([]byte(r.Body))
      //fmt.Printf("%+v\n", r)
      DumpRequest(w, r)
    })

     http.ListenAndServeTLS(":3000", "certs/cert.pem", "certs/key.pem", nil)
}

// DumpRequest is used to dump an incomming request to CLI. This is helpful
// for debugging REST calls.
func DumpRequest(w http.ResponseWriter, req *http.Request) {
    requestDump, err := httputil.DumpRequest(req, true)
    if err != nil {
	//fmt.Fprint(w, err.Error())
        fmt.Println(string(requestDump))
    } else {
      //fmt.Fprint(w, string(requestDump))
      fmt.Println(string(requestDump))
    }
}
