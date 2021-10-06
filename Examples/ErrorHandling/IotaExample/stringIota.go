package IotaExample

import (
	"fmt"
	"strconv"
)

func Iota() {
	s := strconv.Itoa(66)
	fmt.Println(s)
}

func ReturnValue() {
	n, err := strconv.Atoi("cat")
	fmt.Println("Converted number: ", n)
	fmt.Println("Returned error values: ", err)
}

func ErrorHandling() {
	age := "duck"

	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Success: Converted %q to %d.\n", age, n) // happy path
}
