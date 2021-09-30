package main

import "fmt"

func main() {

	color, color2 := "teal", "kentucky red"
	fmt.Println("Original Colors:", color, color2)

	color, color2 = "orange", "green"
	fmt.Println("New colors:", color, color2)

	num1, num2 := 200, 13
	fmt.Println("Original Numbers:", num1, num2)

	num1, num2 = 13, 200
	fmt.Println("New Numbers:", num1, num2)

}
