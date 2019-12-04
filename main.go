package main

import (
	"net/http"

	"github.com/davidzech/webtutorial/notes"
	"github.com/gorilla/mux"
)

// Define a REST API for managing text notes
// What is our resource? Text based Notes
// We want to follow CRUD
// We want to define 4 different endpoints (URL's) to manage the life cycle of notes
// Create -> create a new note
// Read -> fetch an existing note
// Update -> update an existing note
// Delete -> delete a note

var globalDB = notes.NewNotesDB()

// HTTP POST
func handleCreate(w http.ResponseWriter, req *http.Request) {

}

// HTTP GET
func handleRead(w http.ResponseWriter, req *http.Request) {

}

// HTTP PUT
func handleUpdate(w http.ResponseWriter, req *http.Request) {

}

// HTTP DELETE
func handleDelete(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	http.ListenAndServe("", router)
}
