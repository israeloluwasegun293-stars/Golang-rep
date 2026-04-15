package main

import (
	"fmt"
)

func main() {

	// var pointer *int
	// fmt.Println("The value of pointer is:", pointer)

	myNumber := 345

	var ptr = &myNumber

	fmt.Println("The actual value of pointer is;", ptr)
	fmt.Println("the actual value of pointer is:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value is:", myNumber)

}
