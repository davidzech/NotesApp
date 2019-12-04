package notes

// NotesDB is a struct
// Go doesnt have classes, it just has plain data structures
// And you can define methods associated with a struct
// its called a struct rather than a class, cuz go has no inheritance whatsoever
type NotesDB struct {
	notes map[string]string // store notes by name
}

// NewNotesDB creates a new notes db
func NewNotesDB() *NotesDB {
	instance := NotesDB{
		notes: make(map[string]string),
	}

	// var copy NotesDB = instance  // creates a copy of instance
	var ref *NotesDB = &instance // creates a variable that points to the original "instance"

	return ref // return the address of instance, basically meaning return the pointer that points to instance
}

// Add adds a note named "name" with the contents "note" to the *NotesDB object "db"
func (db *NotesDB) Add(name string, note string) {
	db.notes[name] = note
}

func (db *NotesDB) Update(name string, note string) {
	db.notes[name] = note
}

func (db *NotesDB) Read(name string) string {
	return db.notes[name]
}

func (db *NotesDB) Delete(name string) {
	delete(db.notes, name)
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

func test() {
	notesDB := NewNotesDB()
	notesDB.Add("ournote", "i love you")
}
