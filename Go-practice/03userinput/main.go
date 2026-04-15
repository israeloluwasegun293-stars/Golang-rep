package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcome string = "welcome to our IJhayclassics music App"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our App")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string", err)
		return
	}
	fmt.Println("Thanks for rating our App:", input)
	fmt.Printf("variable is of type %T", input)
}