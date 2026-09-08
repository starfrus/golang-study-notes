package notes

import (
	"fmt"
	"notesapp/internal/input"
)

type Note struct {
	ID   int
	Text string
}

func NewNotes(text string) Note {
	return Note{
		Text: text,
	}
}

func NewNotesMust() (Note, error) {
	fmt.Printf("> Введите задачу: ")
	title, err := input.InputData()
	if err != nil {
		return Note{}, err
	}
	if err := input.CheckEmptyInputData(title); err != nil {
		return Note{}, err
	}
	return NewNotes(title), nil
}
