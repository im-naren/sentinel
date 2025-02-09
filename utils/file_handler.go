package utils

// create a file handler to read and write files form csv
import (
	"encoding/csv"
	"fmt"
	"os"
)

// FileHandler struct
type FileHandler struct {
	
}

// NewFileHandler creates a new file handler
func NewFileHandler() *FileHandler {
	return &FileHandler{}
}


// ReadCSV reads a csv file and returns the data
func (f *FileHandler) ReadCSV(filePath string) ([][]string, error) {
	// Open the file
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer csvFile.Close()

	// Read the file
	reader := csv.NewReader(csvFile)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}

// WriteCSV writes data to a csv file
func (f *FileHandler) WriteCSV(filePath string, data [][]string) error {
	// Create the file
	csvFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer csvFile.Close()

	// Write the data
	writer := csv.NewWriter(csvFile)
	err = writer.WriteAll(data)
	if err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}

