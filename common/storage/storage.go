package storage

import (
	"notesapp/common/notes"
)

type Storage interface {
	Add(n notes.Note) error
	List() ([]notes.Note, error)
	Delete(id int) error
}

func AddNote(storage Storage, note notes.Note) error {
	return storage.Add(note)
}

func ListNotes(storage Storage) ([]notes.Note, error) {
	return storage.List()
}

func DeleteNote(storage Storage, id int) error {
	return storage.Delete(id)
}
