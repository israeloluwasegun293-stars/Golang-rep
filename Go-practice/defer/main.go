package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	MyDefer()
}

func MyDefer() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
