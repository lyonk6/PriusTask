package main

import (
     "fmt"
     "net/http"
     "net/http/httputil"
)

func main(){
     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	 w.Write([]byte("\"Welcome to PriusTask!\""))
	 // w.Write([]byte(r.Body))
         //fmt.Printf("%+v\n", r)
         //         fmt.Printf("Here is the body:%T ", r.Body,  r.Body, "\n")
         DumpRequest(w, r)
        	   
    })

     http.ListenAndServeTLS(":3000", "certs/cert.pem", "certs/key.pem", nil)
}

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
