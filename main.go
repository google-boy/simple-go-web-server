package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var  counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang server")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello from Golang server")
	})
	
	fmt.Println("Server running running at https://localhost:8443")
	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil))
	
}