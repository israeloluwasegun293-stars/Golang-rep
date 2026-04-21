package main 

// Divide attempts to divide a by b.
// It returns the quotient and a boolean indicating success.
import "fmt"

func Divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	result, ok := Divide(10, 2)
	if ok {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Cannot divide by zero")
	}
}
