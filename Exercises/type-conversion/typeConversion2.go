package main

import "fmt"

func main() {
	num1, num2 := 10, 5.5
	num1 = int(num2)
	fmt.Println(float64(num1) + num2)

}
