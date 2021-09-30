package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("css/main.css") // Using the Split import
	fmt.Println("dir :",dir)
	fmt.Println("file :", file)
}

/*
'Split' splits path immediately following the final slash, separating it into a directory and file name component.
*/
