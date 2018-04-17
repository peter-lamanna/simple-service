package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//write welcome message to the response buffer
	fmt.Fprintf(w, "Welcome to Magalix Simple Service!! Your Path is: %s!", r.URL.Path[1:])

	//log the call to the standard I/O. You can read these logs at Application web console
	log.Println(fmt.Sprintf("request recieved - path: %s", r.URL.Path))
}

func main() {
	
	//http request handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
	
	//listen to port 8080
	http.ListenAndServe(":8080", nil)
}

//Message is the http parameter
type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {

	//send json text back when user calls the /about
	m := Message{"Welcome to the Simple Service API, v1.0"}
	log.Println("request recieve - path: about")
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}