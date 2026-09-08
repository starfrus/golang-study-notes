package memory

import (
	"errors"
	"fmt"
	"notesapp/internal/notes"
)

type MemoryStorage struct {
	notes  []notes.Note
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		notes: make([]notes.Note, 0),
	}
}

func (m *MemoryStorage) Add(n notes.Note) error {
	n.ID = m.nextID + 1
	m.nextID++
	m.notes = append(m.notes, n)
	return nil
}

func (m *MemoryStorage) List() ([]notes.Note, error) {
	if len(m.notes) == 0 {
		err := errors.New("empty lists notes")
		return []notes.Note{}, fmt.Errorf("! Список задач пуст: %w", err)
	}
	return m.notes, nil
}

func (m *MemoryStorage) Delete(id int) error {

	if id == 0 {
		err := errors.New("id is equal to zero")
		return fmt.Errorf("! Вы указали нулевой id: %w", err)
	}

	if id < 0 {
		err := errors.New("id is negative")
		return fmt.Errorf("! Вы указали отрицательный id: %w", err)
	}

	for i, v := range m.notes {
		if v.ID == id {
			m.notes = append(m.notes[:i], m.notes[i+1:]...)
			fmt.Printf("> Заметка %s удалена", v.Text)
			return nil
		}
	}
	err := errors.New("notes not in memory storage")

	return fmt.Errorf("! Заметка с id:%d :%w", id, err)
}
