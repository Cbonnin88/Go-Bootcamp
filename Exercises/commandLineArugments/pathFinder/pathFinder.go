package main

import (
	"fmt"
	"os"
)

func main() {

	myprogram := os.Args

	fmt.Println("Path:", myprogram[0])

}
