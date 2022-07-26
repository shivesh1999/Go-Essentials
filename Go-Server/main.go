package main

import (
	"fmt"
	"log"
	"net/http"
)

//An http.ResponseWriter which is where you write your text/html response to.
//An http.Request which contains all information about this HTTP request including
//things like the URL or header fields
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error :", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "POST request Successful")
	name := r.FormValue("name")
	fmt.Fprintf(w, " %s\n", name)
}

func main() {
	// FileServer returns a handler that serves HTTP requests
	// with the contents of the file system rooted at root.
	fileServer := http.FileServer(http.Dir("./static"))
	// Handle registers the handler for the given pattern in the DefaultServeMux.
	http.Handle("/", fileServer)
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server at port 8080")
	//Defining the route
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
