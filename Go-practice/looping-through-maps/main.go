package main

import "fmt"

func main() {
	Names := make(map[int]string)

	Names[1] = "Adeola"

	Names[2] = "Olayiwola"

	Names[3] = "Adebimpe"

	for i := 0; i < len(Names); i++ {
		fmt.Println(Names[i])
	}
}
