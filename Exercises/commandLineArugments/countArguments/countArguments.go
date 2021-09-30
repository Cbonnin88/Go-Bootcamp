package main

import (
	"fmt"
	"os"
)

func main() {

	count := len(os.Args) - 1

	fmt.Println("1st Argument:", os.Args[1])
	fmt.Println("2nd Argument:", os.Args[2])
	fmt.Println("3rd Argument:", os.Args[3])
	fmt.Println("4th Argument", os.Args[4])

	fmt.Printf("There are %d names.\n", count)
}
