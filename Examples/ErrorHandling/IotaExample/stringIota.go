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
