package notes

import "errors"

// NotesDB is a struct
// Go doesnt have classes, it just has plain data structures
// And you can define methods associated with a struct
// its called a struct rather than a class, cuz go has no inheritance whatsoever
type NotesDB struct {
	notes map[int]string // store notes by name
}

// NewNotesDB creates a new notes db
func NewNotesDB() *NotesDB {
	instance := NotesDB{
		notes: make(map[int]string),
	}

	// var copy NotesDB = instance  // creates a copy of instance
	var ref *NotesDB = &instance // creates a variable that points to the original "instance"

	return ref // return the address of instance, basically meaning return the pointer that points to instance
}

var globalCounter int = 0

// Add adds a note named "name" with the contents "note" to the *NotesDB object "db
// Returns a nil error on success, or an error on failure
func (db *NotesDB) Add(note string) (int, error) {
	globalCounter++
	db.notes[globalCounter] = note
	return globalCounter, nil
}

func (db *NotesDB) Update(id int, note string) {
	db.notes[id] = note
}

func (db *NotesDB) Read(id int) (string, bool) {
	val, exists := db.notes[id]
	return val, exists
}

type Note struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func (db *NotesDB) Find() []Note {
	out := []Note{}
	for k, v := range db.notes {
		out = append(out, Note{
			ID:    k,
			Value: v,
		})
	}
	return out
}

func (db *NotesDB) FindBy(criteria string) []Note {
	return nil
}

func (db *NotesDB) Delete(id int) error {
	if _, exists := db.notes[id]; exists == false {
		return errors.New("note with id " + string(id) + " does not exist")
	}
	delete(db.notes, id)
	return nil
}

// in python
/*
class NotesDB

def __init__():
	self.notes = dict()

def add(self, name, note):
	self.notes[name] = note

db = NotesDB()
db.Add("ournote", "i love you")

*/
