package main

import (
	"fmt"
)

const LoginToken string = "ishdfjvidijvn"

func main() {
	var username string = "Israel"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.990349780349
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default Value and aliases for int
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//default value for string
	var stringVariable string
	fmt.Println(stringVariable)
	fmt.Printf("variable is of type: %T \n", stringVariable)

	//default value for bool
	var boolValue bool
	fmt.Println(boolValue)
	fmt.Printf("variable is of type: %T \n", boolValue)

	//implicit type
	var website = "222.5556"
	fmt.Println(website)
	fmt.Printf("variable is of: %T \n", website)

	//the no var style
	numberOfUsers := 3000000000000
	fmt.Println(numberOfUsers)
	fmt.Printf("variable is of type: %T \n", numberOfUsers)

	//for const|
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

	//printf (formattefd string)
	Age := 22
	name := "Israel"

	//%v to print the first variable
	fmt.Printf("my age is %v and my name is %v \n", Age, name)

	//%Q to add quotes
	fmt.Printf("my age is %v and my name is %v \n", Age, name)

	//%T to print variable type
	fmt.Printf("age is of type %T", Age)

	//%f for Float variable indication
	fmt.Printf("You scored %f in your exam! \n", 222.567)

	//%...f to indicate how many decimal numbers you want
	fmt.Printf("You scored %0.2f in your exam \n", 234.576389874856)

	//Sprintf (saved formatted string)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", Age, name)
	fmt.Println("The saved string is:", str)

}
