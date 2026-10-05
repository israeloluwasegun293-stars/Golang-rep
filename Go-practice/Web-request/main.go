package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://roadmap.sh/ai/course/go-programming-from-fundamentals-to-web-services?fl=0"

func main() {

	fmt.Println("Israel is progressing, keep it up!")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("response if of type: %T\n", response)
	}
	defer response.Body.Close() //Is it the responsibility of the user to close the connection after making a request!!!

	readBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(readBytes))
}
