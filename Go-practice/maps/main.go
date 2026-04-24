package main

import "fmt"

func main() {
	myNames := make(map[string]string)

	myNames["Israel"] = "Oluwasegun"

	myNames["Okunade"] = "Joshua"

	myNames["Oluwaseyi"] = "JesuTofunmi"

	fmt.Println(myNames, len(myNames))
	fmt.Println(myNames["Okunade"], len(myNames))
	delete(myNames, "Oluwaseyi")
	fmt.Println(myNames["Okunade"], len(myNames))
}
