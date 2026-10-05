package main

import "fmt"


type Player struct {
	Name string
	Age int
}

func (p Player) levelUp() {
	p.Name += "Israel"
}


func main() {

	// var pointer *int
	// fmt.Println("The value of pointer is:", pointer)

	// myNumber := 345

	// var ptr = &myNumber

	// fmt.Println("The actual value of pointer is;", ptr)
	// fmt.Println("the actual value of pointer is:", *ptr)

	// *ptr = *ptr * 2
	// fmt.Println("new value is:", myNumber)

	// score := 2345

	// ptr2 := &score
	// fmt.Println(score)
	// fmt.Println("the address of the variable is :", ptr2)
	// fmt.Println("the original value held in the pointer is :", *ptr2)
	number := 100
	fmt.Println(passBvalue(&number))

	myDee := Player{Name: "Okunade", Age: 22}
	levelUp(myDee)
	fmt.Println(myDee.Name)
}

func passBvalue(num *int) int {
	*num = *num + *num
	return *num
}

// When you pass variables to functions in Go, the language uses "pass by value."
// This means Go creates a copy of the variable for the function to use.
// If you want a function to modify the original variable rather than a copy, you must pass a pointer.

//REAL WORLD INSTANCES OF THE POINTER
