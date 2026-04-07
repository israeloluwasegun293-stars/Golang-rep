//this codd takes the arguement from a file (input.txt)

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	inputFile := os.Args[1]

// 	num, err := os.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Println("error reading file", err)
// 		return
// 	}
// 	data := string(num)

// 	datas := strings.TrimSpace(data)

// 	dat, err := strconv.ParseInt(datas, 10, 64)
// 	if err != nil {
// 		fmt.Println("error converting", err)
// 		return
// 	}
// 	intdata := float64(dat)

// 	solution := intdata*9.0/5.0 + 32

// 	fmt.Println(solution)
// }

//this code takes this arguement direectly from  the terminal

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguements")
		return
	}

	num := os.Args[1]

	datas := strings.TrimSpace(num)

	dat, err := strconv.ParseInt(datas, 10, 64)
	if err != nil {
		fmt.Println("error converting", err)
		return
	}
	intdata := float64(dat)

	solution := intdata*9.0/5.0 + 32

	fmt.Println(solution)
}
