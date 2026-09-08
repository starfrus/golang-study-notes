package memory

import (
	"errors"
	"fmt"
	"notesapp/common/notes"
)

type MemoryStorage struct {
	notes  []notes.Note
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		notes:  make([]notes.Note, 0),
		nextID: 1,
	}
}

func (m *MemoryStorage) Add(n notes.Note) error {
	n.ID = m.nextID
	m.nextID++
	m.notes = append(m.notes, n)
	return nil
}

func (m *MemoryStorage) List() ([]notes.Note, error) {

	return m.notes, nil
}

func (m *MemoryStorage) Delete(id int) error {

	for i, v := range m.notes {
		if v.ID == id {
			m.notes = append(m.notes[:i], m.notes[i+1:]...)
			fmt.Printf("> Заметка [%s] удалена\n", v.Text)
			return nil
		}
	}
	err := errors.New("notes not in memory storage")

	return fmt.Errorf("! Заметка с id: %d отсутствует: %w", id, err)
}
