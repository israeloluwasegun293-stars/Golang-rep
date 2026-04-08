package main

import "fmt"


func main() {

	os := "Linux"

	switch os{

	case "Linux":
		fmt.Println("operating system is Linux")

	case "Windows":
		fmt.Println("Operating system is Windows")

	default:
		fmt.Println("operating system not recognized")
	}
}