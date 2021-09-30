package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	msg := `Hi ` + name + ` how are you?`

	fmt.Println(msg)

}
