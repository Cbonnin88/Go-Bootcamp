package main

import "fmt"

func main() {

	color := "green"
	fmt.Println("Original Color:", color)

	color = "dark " + color
	fmt.Println("New Color:", color)

}
