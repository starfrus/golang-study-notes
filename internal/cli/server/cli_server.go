package cli_server

import (
	"fmt"
	"notesapp/common/notes"
	"notesapp/common/storage"
	"notesapp/internal/cli/input"
	"notesapp/internal/cli/memory"
	"os"
)

func HandleCommand(s storage.Storage, command string) {
	switch command {
	case "add":
		note, err := notes.NewNotesMust()
		if err != nil {
			fmt.Println(err)
		} else {
			storage.AddNote(s, note)
			fmt.Println("> Задача добавлена!")
		}
	case "list":
		notes, err := storage.ListNotes(s)
		if err != nil {
			fmt.Println(err)
		} else if len(notes) == 0 {
			fmt.Println("! Список задач пуст")
		} else {
			for _, v := range notes {
				fmt.Printf("> ID: %3d\t Title: %s\n", v.ID, v.Text)
			}
		}
	case "del":
		id, err := input.InputID()
		if err != nil {
			fmt.Println(err)
		} else {
			if err := storage.DeleteNote(s, id); err != nil {
				fmt.Println(err)
			}
		}
	case "help":
		fmt.Println()
		fmt.Println("add  -  добавить задачу")
		fmt.Println("list -  получить список задач")
		fmt.Println("del  -  удалить задачу")
		fmt.Println("quit -  выход из программы")
		fmt.Println()
	case "quit":
		fmt.Println("> Пока:)")
		os.Exit(0)
	default:
		fmt.Println("> Неизвестная команда, напиши help")
	}
}

func StartCLI() {
	ms := memory.NewMemoryStorage()
	for {
		fmt.Printf("> Выберите команду [add/list/del/help/quit]: ")
		command, err := input.InputData()
		if err != nil {
			panic(err)
		}
		HandleCommand(ms, command)
	}
}
