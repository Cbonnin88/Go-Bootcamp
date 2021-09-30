package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("secret/file.txt")

	fmt.Println("Directory:",dir)
	fmt.Println("file:",file)

}
