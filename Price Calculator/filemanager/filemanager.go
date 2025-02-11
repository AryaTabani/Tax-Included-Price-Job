package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line in file.")
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file")
	}
	encoder := json.NewDecoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to convert data to json")
	}
	file.Close()
	return nil
}
func New(inputpath, outputpath string) FileManager {
	return FileManager{
		InputFilePath:  inputpath,
		OutputFilePath: outputpath,
	}
}
