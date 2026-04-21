package main

import "fmt"

func main() {

	sum := 0

	stopppedat := 0

	for i := 1; i <= 100; i++ {

		sum += i
		stopppedat = i

		if sum > 500 {
			break
		}
	}
	fmt.Printf("The sum of all numbers is %v \n", sum)
	fmt.Printf(" at: %v \n", stopppedat)
}
