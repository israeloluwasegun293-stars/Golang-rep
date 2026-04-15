package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go time study")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:09:02 Monday"))

	createdTime := time.Date(2026, time.April, 15, 20, 03, 0, 0, time.UTC)

	fmt.Println(createdTime)
}
