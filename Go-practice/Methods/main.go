package main

import "fmt"

func main() {

	Israel := User{"Israel", "israeloluwasegun293@gmail.com", true, 19}
	fmt.Println(Israel)
	Israel.GetStatus()
	Israel.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "israeljoshua125@gmail.com"
	fmt.Println("this user's Email is :", u.Email)
}
