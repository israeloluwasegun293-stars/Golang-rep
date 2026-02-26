package main

import "fmt"

func main() {
	var names = [4]int{234, 345, 456, 567}

	var ages = [3]string{"Jude", "ola", "Jube"}

	var years = [4]int{2000, 2001, 2002, 2003}

	fmt.Println(names, len(names))
	fmt.Println(ages, len(ages))
	fmt.Println(years, len(years))

	//slices this time around

	var Uni = []string{"Uniosun", "Unilag", "Lasu"}
	Uni[0] = "Kwasu"
	Uni = append(Uni, "Lasustech")

	fmt.Println(Uni, len(Uni))

	//slice ranges
	rangeOne := Uni[0:3]
	rangeTwo := Uni[:4]
	rangeThree := Uni[0:]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	//we can also append in slices ranger
	rangeThree = append(rangeThree, "IJhay")

	fmt.Println(rangeThree)
}
