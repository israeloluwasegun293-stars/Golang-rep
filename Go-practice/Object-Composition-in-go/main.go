package main

import "fmt"

func main() {

	aCar := car{}
	fmt.Println(aCar)
	aCar.mpg = "234"
	(aCar.getMpg())
}

type vehicle struct {
	mpg string
}

type car struct {
	vehicle
}

func (v *vehicle) getMpg() {
	fmt.Println("the value of the Mpg is", v.mpg)
}
