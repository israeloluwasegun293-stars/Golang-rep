package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	//rule for number of Argument
	if len(os.Args) != 3 {
		fmt.Println("Invalid number of Arguments")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	Data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	text := string(Data)

	isreal := text + "!!!"

	err = os.WriteFile(outputFile, []byte(isreal), 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}
	fmt.Println("Succesful!!! Check result.txt")

}

func ReadFile(filepath string) ([]string, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	fileStat, err2 := file.Stat()
	if err2 != nil {
		return nil, err2
	}

	fileBuffer := make([]byte, fileStat.Size())

	_, err3 := file.Read(fileBuffer)
	if err3 != nil {
		return nil, err3
	}

	if len(fileBuffer) == 0 {
		return nil, errors.New("File is empty")
	}

	output := strings.Fields(string(fileBuffer))

	return output, nil
}

func WriteFile(filepath string, content string) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err2 := file.WriteString(content)
	if err2 != nil {
		return err
	}
	return nil
}
