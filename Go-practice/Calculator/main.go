package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please provide input1")

	input, _ := reader.ReadString('\n')

	num1, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Please provide a valid input")
		return
	}

	fmt.Println("Please provide the second input")
	input2, _ := reader.ReadString('\n')

	num2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println("Please provide a valid input")
		return
	}

	fmt.Println("Please provide a logic operator(+, -, /, *)")
	input3, _ := reader.ReadString('\n')

	operator := strings.TrimSpace(input3)

	switch operator {
	case "+":

		fmt.Printf("Result: %.2f\n", num1+num2)

	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)

	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("cannot divide number by 0")
			return
		}
		fmt.Printf("Result: %.2f\n", num1/num2)
	default:
		fmt.Println("Operator not recognized, Please provide a valid operator")
	}

}
