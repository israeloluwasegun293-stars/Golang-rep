package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
}

type Robot struct {
	Name  string
	Model string
}

type Behavior interface {
	walk() string
	talk() string
	sing() string
	playKeyboard() bool
}

func (p *Person) walk() string {
	return "Walks normal like human being"
}

func (p *Person) talk() string {
	return "talks like audio"
}

func (p *Person) sing() string {
	return "sings beautifully"
}

func (p *Person) playKeyboard() bool {
	return true
}

func (r *Robot) walk() string {
	return "rolls on wheel"
}

func (r *Robot) talk() string {
	return "talks like audio"
}

func (r *Robot) sing() string {
	return "Plays audio files"
}

func (r *Robot) playKeyboard() bool {
	return true
}



func describe(b Behavior){
	fmt.Println(b.walk())
	fmt.Println(b.talk())
	fmt.Println(b.sing())
	fmt.Println(b.playKeyboard())
}
func main() {
	p := &Person{Firstname: "IJhayclassics"}
	r := &Robot{Name: "Jarvis"}

	describe(p)
	describe(r)
}
