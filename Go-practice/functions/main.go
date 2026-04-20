package main

import "fmt"

func isEven(val int) bool {
	return val%2 == 0
}
func main() {

	numType := isEven(33)

	fmt.Println(numType)
}

/*Create a function named isEven that accepts an integer and returns a boolean (true if the number is even, false otherwise).
Call this function from main with different inputs and print the results.
*/
