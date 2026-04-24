package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

// type behaviour interface {
// 	walk()
// 	talk()
// }

func main() {

	p := new(person)
	fmt.Println(p.walk())

	t := new(person)
	fmt.Println(t.talk())
}

func (w *person) walk() string {
	return "Have a good day"

}

func (w *person) talk() string {
	return "I am a Boy"
}
