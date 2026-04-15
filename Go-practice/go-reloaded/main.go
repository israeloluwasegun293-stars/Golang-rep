package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid number of arguement")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error readingfile:")
		return
	}

	data := string(file)

	arrayData := strings.Fields(data)

	joinedData := strings.Join(arrayData, " ")

	result := processText(joinedData)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
	}
	fmt.Println("sucessfully read file!")
}
func processText(text string) string {
	return text
}
// 08159687744