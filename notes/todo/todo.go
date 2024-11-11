package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t Todo) Display() {
	jsonData, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}

func (t Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.Marshal(t)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{Text: content}, nil
}
