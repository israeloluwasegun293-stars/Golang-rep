package main

import "fmt"

func main() {
	// inventory := make(map[string]int)

	// inventory["Laptop"] = 30
	// inventory["phones"] = 20
	// inventory["ipads"] = 12
	// inventory["samsung"] = 50

	// fmt.Printf("the number of laptops in the shop is %v\n", inventory["Laptop"])
	// fmt.Println(thisMap())

	// check()
	mapsStruct()

	// values := []Details{
	// 	{Name:"Israel", Age: 22, Status: true},
	// }
	// fmt.Println(values)
}

// type Details struct {
// 	Name string
// 	Age int
// 	Status bool
// }

// //this is how maps of struct actually works; we use the struct to select the parameters we're passing in and the type the use the map to put in the values

// func check() {
// checkout := make(map[int]Details)

// checkout[1] = Details{Name: "Israel", Age: 22, Status: true}
// checkout[2] = Details{Name: "Esther", Age: 18, Status: true}
// checkout[3] = Details{Name: "Abraham", Age: 90, Status: false}

// fmt.Println(checkout)
// }

type Details struct {
	Name		string
	Age        int
	Status     bool
	Location   string
	Occupation string
}

func mapsStruct() {
	theWork := make(map[int]Details)

	theWork[1] = Details{Name: "Akolade Ademola,", Age: 44, Status: true, Location: "Osun state,", Occupation: "Actor"}
	theWork[2] = Details{Name: "Bimpe Bobola", Age: 30, Status: true, Location: "Ogun state", Occupation: "Developer"}
	theWork[3] = Details{Name: "Jekolo olamide", Age: 39, Status: true, Location: "Ondo state", Occupation: "chef"}
	fmt.Println(theWork)
}


// func thisMap()(map[string]int) {
// 	myMap := make(map[string]int)

// 	myMap["Israel"] = 22
// 	myMap["Esther"] = 18

// 	comma, ok := myMap["Israel"]
// 	if ok {
// 		fmt.Println(comma, ok)
// 	}

// 	comma, ok = myMap["taiwo"]
// 	if !ok {
// 		fmt.Println(comma, ok)
// 	}

// 	for i, char := range myMap {
// 		fmt.Printf("his name is %s, and his age is %d\n", i, char)
// 	}
// 	delete(myMap, "Israel")

// 	myMap["Esther"] = 20

// 	return myMap
// }