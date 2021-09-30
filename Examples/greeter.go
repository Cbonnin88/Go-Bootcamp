package main

import (
	"fmt"
	"os" // Allow me to access the operating system functionalities
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path:", os.Args[0])
}
