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
	file, err := os.Open(fm.InputPath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() { // move uma linha para frente
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		file.Close()
		return nil, err
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteLines(data any) error {
	file, err := os.Create(fm.OutputPath)
	if err != nil {
		return errors.New("erro ao criar o novo arquivo")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("erro ao converter dados para JSON")
	}

	file.Close()
	return nil
}

func New(inputPath string, outputPath string) *FileManager {
	return &FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}
