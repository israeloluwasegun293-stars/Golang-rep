package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("Running on Linux")
	case "windows":
		fmt.Println("Running on Windows")
	case "darwin":
		fmt.Println("Running on macOS")
	default:
		fmt.Println("Unknown OS")
	}
}
