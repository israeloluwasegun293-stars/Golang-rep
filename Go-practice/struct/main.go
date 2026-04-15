// package main

// import "fmt"

// type user struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	// method one
// 	data := user{"Israel", "Ikorodu", 20}
// 	fmt.Printf("My name  is %s \nI currently reside in %s \nI am %d years old\n",
// 		data.name,
// 		data.city,
// 		data.age,
// 	)
// 	//method 2
// 	user1 := user{
// 		name: "Israel",
// 		city: "ikorodu",
// 		age:  20,
// 	}
// 	fmt.Printf("My name is %s \nI live in %s \nI am %d years old\n",

// 		user1.name,
// 		user1.city,
// 		user1.age,
// 	)
// 	//3rd method
// 	var user2 user

// 	user2.name = "Israel"

// 	user2.city = "Ikorodu"

// 	user2.age = 20

// 	fmt.Printf("My name is %s \nI stay in %s \nI am %d years old\n",
// 		user2.name,
// 		user2.city,
// 		user2.age,
// 	)
// 	//we can also modify the sruct but this won't apply cuz it's before the fmt.print
// 	user2.age = 100
// }

package main

import (
	"fmt" 
)

type Student struct {
	Name string

	Age int

	Score int
}

func main() {

	students := []Student{

		{Name: "Ade", Age: 19, Score: 54},
		{Name: "Scott", Age: 23, Score: 45},
		{Name: "Jude", Age: 12, Score: 34},
		{Name: "Ola", Age: 23, Score: 90},
		{Name: "Williams", Age: 32, Score: 70},
		{Name: "James", Age: 24, Score: 49},
	}

	for _, students := range students {
		fmt.Printf("Student name is: %s\n, Student is %d years old\n, score is: %d\n",

			students.Name,
			students.Age,
			students.Score,
		)
	}
}
