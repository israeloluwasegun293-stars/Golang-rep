package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid number of arguements")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	data := string(file)

	arrayData := strings.Fields(data)

	joinedArray := strings.Join(arrayData, " ")

	err = os.WriteFile(outputFile, []byte(joinedArray), 0644)
	if err != nil {
		fmt.Println("Error writing file")
	}
	fmt.Println("sucessfully written file")
}
