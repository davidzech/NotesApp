// +build ignore

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // mux is a framework that lets us do cool URL handling for a web server
)

// write a web page such that http://localhost/hello/<your name here> give us a web page that says hello to you!
func helloHandler(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(200)
	var variableMap map[string]string = mux.Vars(req) // mux.Vars examines req object, and matches req.URL with the pattern "/hello/{name}" that we registered below
	name := variableMap["name"]
	// name := mux.Vars(req)["name"]
	htmlStr := fmt.Sprintf(`<html>Hello %s<br/>URL: %s</html>`, name, req.URL.String())

	w.Write([]byte(htmlStr))
}

// we want to do /hello/<firstname>/<lastname>

// Browser initiates a connection to the server -> (http://localhost/hello/anh/nguyen) -> packaged into a HTTP request object (req)
// Server which is listening for connections, obtains an HTTP request from the browser
// Browser (Request) -> Server
// our code is this point inbetween
// Broser <- (Response) Server
// The http server parses this request, does its internal logic, and then sends back a HTTP Response
// HTTP Response (StatusCode, Body)
func helloFullName(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req) // does pattern matching magic of the req.URL for us
	firstName := vars["firstname"]
	lastName := vars["lastname"]

	htmlStr := fmt.Sprintf("Hello %s %s", firstName, lastName)

	// Below two items make up the basics of an HTTP Response
	w.WriteHeader(200)       // Status Code
	w.Write([]byte(htmlStr)) // Body
}

// Keywords:
// REST API
// CRUD (Create, Read, Update, Delete) -> for a resource, define 4 endpoints to create, read, update, delete
// Philosophies for creating backend web api's

func main() {
	router := mux.NewRouter()

	// say we visit /hello/anh, the router will match that pattern with /hello/{name} and extract "anh" into the "name" variable in its dictionary
	router.HandleFunc("/hello/{name}", helloHandler) // match URL paths for /hello/<some text here> and store that text in a dictionary with key "name"
	router.HandleFunc("/hello/{firstname}/{lastname}", helloFullName)
	http.ListenAndServe(":80", router)
}
