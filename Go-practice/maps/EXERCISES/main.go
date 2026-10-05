package main

import "fmt"

func main() {
	Tracker := make(map[string]float64)

	Tracker["Coaster-Buscuit"] = 50.00
	Tracker["Short-bread"] = 350.00
	Tracker["Yamaha-keyboard-Psr-Sx750"] = 1500500.00
	Tracker["Kaysen Guitar"] = 225000.00

	for i, num := range Tracker {
		fmt.Printf("The item price is: %s and the price is: %f \n", i, num)
	}
	slice := []string{"pawpaw", "mango","pawpaw", "mango", "mango", "mango", "Apple","Apple","Apple","Apple"}
	fmt.Println(Frequency(slice))
	safeDel("ipads")

}

func Frequency(s []string) map[string]int {

	counts := make(map[string]int)
	for _, num := range s {
		counts[num]++
	}
	return counts
}

func safeDel(s string) {
	inventory := make(map[string]int)

	inventory["Laptop"] = 30
	inventory["phones"] = 20
	inventory["ipads"] = 12
	inventory["samsung"] = 50

	_, ok := inventory[s]
	if ok {
		delete(inventory, s)
		fmt.Println("Item has been found and sucessfully deleted")
	}else{
	fmt.Println("item was not found")
	//3.....Safe Deletion: Write a code snippet that attempts to delete a key from your inventory map, 
// but prints a confirmation message only if the key actually existed before the deletion.
	}

}

//1...Inventory Tracker: Create a map to store product names (string) and their prices (float64). Add three items to the map.
// Write a loop to print out each item and its price.

//2.....Frequency Counter: Write a function that takes a slice of strings (e.g., []string{"apple", "orange", "apple", "banana"}) 
// and returns a map where the keys are the fruit names and the values are the counts of how many times each fruit appears in the slice.

//3.....Safe Deletion: Write a code snippet that attempts to delete a key from your inventory map, 
// but prints a confirmation message only if the key actually existed before the deletion.
