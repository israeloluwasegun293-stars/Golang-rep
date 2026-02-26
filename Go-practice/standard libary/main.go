package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "Welcome to our Music app"
	//this will check is  that word we input is actually in the string

	fmt.Println(strings.Contains(greetings, "app"))

	//this replaces a word in the string
	fmt.Println(strings.ReplaceAll(greetings, "app", "application"))

	//this converts the whole string to UpperCase
	fmt.Println(strings.ToUpper(greetings))

	//index gives us the number of the position a certain letter is
	fmt.Println(strings.Index(greetings, "ou"))

	//splits the string into slices if we put a particular word inside the " ", it removes that particular word from the slice
	fmt.Println(strings.Split(greetings, " "))

	//the original value itself is unchanged
	fmt.Println("the original value is unchanged:", greetings)

	//sort; it alphabetically arrange the whole thing in the slice

	age := []int{29, 28, 27, 26, 25, 24, 23}

	sort.Ints(age)
	fmt.Println(age)

	index := (sort.SearchInts(age, 26))
	fmt.Println(index)

	//sorts slice of strings
	names := []string{"Zebra", "Dog", "Goat", "Cheeatah"}
	sort.Strings(names)
	fmt.Println(names)

	//to search strings this time
	fmt.Println(sort.SearchStrings(names, "Zebra"))
}
