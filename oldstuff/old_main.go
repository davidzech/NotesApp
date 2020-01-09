// +build ignore

package main // declares that we are adding code to package main

import (
	http "net/http"
)

// w and req are passed to helloHandler by the internal http.Server implementation, that the Go library authors made
func helloHandler(w http.ResponseWriter, req *http.Request) {
	// w is a ResponseWriter, which lets us send data back tgo the browser (aka write data back to the browser)
	// the browser wants to load http://localhost/, it wants an http response back

	// HTTP response has the following:
	// Header
	// Body

	// 404 - means content not found
	// 401 - you arent logged in to access this
	// 200 - means successful (send this when everything is good)
	w.WriteHeader(200)
	w.Write([]byte(`<html>Hello World!</html>`))
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is a home page"))
}

func main() {
	http.HandleFunc("/", helloHandler) // tell the http server to handle at path /, call helloHandler to handle that request
	http.HandleFunc("/home", homeHandler)

	// Protocol called TCP, web browses use TCP (transmission control protocol) to communicate over the internet
	// Client -> Server that is listening on a port
	// Client initiates a connection to a server always, never the other way around
	// Server is always listening for connections on a Port
	// valid ports 1-65535
	// when you type localhost in the browser, it's connecting to http://localhost:80/
	// a server can have like multiple server applications, such as an http server, or an SSH daemon
	// only 1 application at a time, can use a single port.
	// Some common port, standard ports are as such:
	// HTTP always on port 80 by default
	// HTTPS always on port 443 by default
	// SSH is always on port 22 by default
	// http://localhost:8080
	http.ListenAndServe("", nil)
}
