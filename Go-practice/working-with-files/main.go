package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("We're in reading files in Go")

	content := "This needs to go in a file Ijhayclassics"

	file, err := os.Create("My-file.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The length is :", length)

	defer file.Close()
}
