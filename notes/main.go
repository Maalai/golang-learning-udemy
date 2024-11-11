package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {

	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println("unable to save Todo")
		return
	}

	fmt.Println("Todo saved successfully")

	err = outputData(note)

	if err != nil {
		fmt.Println("unable to save notes")
		return
	}

	fmt.Println("notes saved successfully")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	return nil
}

func outputData(outputData outputtable) error {
	outputData.Display()
	return saveData(outputData)
}
