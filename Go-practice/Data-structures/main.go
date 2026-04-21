package main

import "fmt"

func main() {

	// Creating and Initializing Slices

	task := []string{"Create user email", "deploy my website"}

	myTask := make([]int, 12, 456) // takes only 2 or 3 arguements

	fmt.Println(task, "\n", myTask)

	// Manipulating Slice Contents
	var logs []string

	logs = append(logs, "user Login,", "user data,", "User password")
	fmt.Println(logs)

	// Append multiple items at once using the ... variadic operator
	moreLogs := []string{"web services", "domain name", "user's request"}

	logs = append(logs, moreLogs...)
	fmt.Println(logs)

	// Understanding Slicing and Memory
	data := []int{23, 34, 45, 56, 67, 78, 89}

	subset := data[1:5]

	subset[0] = 67

	fmt.Println(data)

	/*Create a slice of integers with a length of 0 and a capacity of 5.
	Loop 5 times to append the numbers 10, 20, 30, 40, 50 to it.
	Print the slice, its length, and its capacity after each append.*/

	myNumbers := []int{}

	ints := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(ints); i++ {
		myNumbers = append(myNumbers, ints[i])

		fmt.Println("Result:", myNumbers)
		fmt.Println(len(myNumbers))
		fmt.Println(cap(myNumbers))
	}
}
