package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	filePtr, err := os.Open(fm.InputPath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(filePtr)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		filePtr.Close()
		return nil, errors.New("failed to read file contents")
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {

	filePtr, err := os.Create(fm.OutputPath)

	if err != nil {
		filePtr.Close()
		return errors.New("error creating file")
	}

	jsonEncoder := json.NewEncoder(filePtr)
	err = jsonEncoder.Encode(data)

	if err != nil {
		filePtr.Close()
		return errors.New("failed to convert data to json")
	}

	filePtr.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{InputPath: inputPath, OutputPath: outputPath}
}
