package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Ijhayclassics Music Worldwide"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating for our brand")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", input, "have a nice time")
	fmt.Printf("rating is of: %T", input)
}
