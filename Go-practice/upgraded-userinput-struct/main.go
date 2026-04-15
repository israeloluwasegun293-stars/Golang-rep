package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rating struct {
	UserName string
	Score    string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Rate our app: ")
	rating, _ := reader.ReadString('\n')

	userRating := Rating{
		UserName: strings.TrimSpace(name),
		Score:    strings.TrimSpace(rating),
	}

	fmt.Printf("\nThanks %s! You rated us: %s\n",
		userRating.UserName,
		userRating.Score,
	)
}