package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// HTTP POST - corresponds to creating an item
// HTTP Post into /notes
// returns a new notes object
func handleCreate(w http.ResponseWriter, req *http.Request) {
	// call into the notes database and create a new one
	decoder := json.NewDecoder(req.Body) // decodes the JSON data from req.Body

	var note notes.Note
	decoder.Decode(&note) // parse req.Body for json data, and populate the fields in n
	id, _ := globalDB.Add(note.Value)
	note.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	encoder.Encode(note)

}

// HTTP GET - corresponds to reading an item
// HTTP Get a note by id at /notes/{id_here} - ex /notes/4
// returns an existing notes object
func handleRead(w http.ResponseWriter, req *http.Request) {
	// fetch the id from the URL pattern
	id, _ := strconv.Atoi(mux.Vars(req)["id"]) // string to int converts a string like "10" to integer 10
	noteContents, exists := globalDB.Read(int(id))
	if exists == false {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found!"))
		return
	}
	note := notes.Note{
		ID:    id,
		Value: noteContents,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(note)
}

// HTTP PUT - corresponds to replacing an item (or creating it if it doesnt exit) - good for updating an item
// HTTP Put a note at an id /notes/{id_here} - ex /notes/4
// Accepts a note object
// Returns a note object with HTTP Status Code 200 (OK)
func handleUpdate(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	decoder := json.NewDecoder(req.Body) // decodes the JSON data from req.Body
	var note notes.Note
	decoder.Decode(&note)

	globalDB.Update(id, note.Value)

	note.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(note)
}

// HTTP DELETE - deletes an item
// HTTP Delete a note at an id /notes/{id_here} - ex /notes/4
// Doesn't accept any body
// Doesn't return anything

func handleDelete(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	err := globalDB.Delete(id)
	if err != nil {
		// there was a problem deleting
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent) // we aren't returning any data, just delete
}

func handleFind(w http.ResponseWriter, req *http.Request) {
	notes := globalDB.Find()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(notes)
}

func main() {
	router := mux.NewRouter()

	// this style of API Design/Architecture
	// is called a REST api
	// Our resource here is a Note
	// and our REST API defines the HTTP endpoints or URLs to manage the full life cycle of managing Note Objects
	router.HandleFunc("/notes", handleCreate).Methods("POST")
	router.HandleFunc("/notes", handleFind).Methods("GET") // Find ALL Notes (with some filter criteria)
	router.HandleFunc("/notes/{id}", handleRead).Methods("GET")
	router.HandleFunc("/notes/{id}", handleUpdate).Methods("PUT")
	router.HandleFunc("/notes/{id}", handleDelete).Methods("DELETE")

	// handle loading the home page
	// serve HTML so we can have a UI at "/"

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	http.ListenAndServe("", router)
}

// backend
// top layer is a HTTP web api (usually a REST API)
// Consumers and Customers and Clients call into our service using REST APIs
// POST - localhost/notes -> a new note object -

// 2 Layers
// REST API on top layer
// --  Database Layer that could be MongoDB, PostgreSQL, MySQL, or our own database implementation
