package main

import "fmt"

func main() {
	inventory := make(map[string]int)

	inventory["Laptop"] = 30
	inventory["phones"] = 20
	inventory["ipads"] = 12
	inventory["samsung"] = 50

	fmt.Printf("the number of laptops in the shop is %v\n", inventory["Laptop"])
}
