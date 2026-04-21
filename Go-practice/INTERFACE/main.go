package main

import "fmt"

type UserInterface interface {
	ActiveUser()
}
type USer1info struct {
	Name   string
	Age    int
	Status bool
}

type User2Info struct {
	Name   string
	Age    int
	City   string
	Status bool
}

func main() {
	var user1 UserInterface
	user1 = USer1info{"Israel", 22, true}
	user1.ActiveUser()
	fmt.Println(user1)
	Israel := USer1info{"Israel", 22, true}

	fmt.Println(Israel.Name)
	// Ola := User2Info{"Ola", 345, "Yaba", true}

}

func (u User2Info) ActiveUser() {

}

func (u USer1info) UserAge() {

}
func (u USer1info) ActiveUser() {

}
