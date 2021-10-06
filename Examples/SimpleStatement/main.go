package main

import (
	"Go-BootCamp/Examples/SimpleStatement/Scope"
	"fmt"
	"strconv"
)

func main() {


	// A simple short if statement
	if n, err := strconv.Atoi("33"); err == nil {
		fmt.Println("No error found, n is", n)
	}

	Scope.Scope()
	Scope.Shadowing()

}
