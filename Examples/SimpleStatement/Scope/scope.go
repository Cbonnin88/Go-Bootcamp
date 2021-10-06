package Scope

import (
	"fmt"
	"os"
	"strconv"
)

func Scope() {
	if a := os.Args; len(a) != 2 {
		// only: a variable
		fmt.Println("Give me a number")

	} else if n, err := strconv.Atoi(a[1]); err != nil {

		fmt.Printf("Cannot cnvert %q.\n", a[1])

	} else {

		fmt.Printf("%s * 2 = %d\n",a[1], n*2)
	}
}

func Shadowing() {
	var (
		n int
		err error
	)

	if a := os.Args; len(a) != 2 {

		fmt.Println("Give me a number")

	} else if n, err = strconv.Atoi(a[1]); err != nil {

		fmt.Printf("Cannot cnvert %q.\n", a[1])

	} else {
		n *=2
		fmt.Printf("%s * 2 = %d\n",a[1], n)
	}
	fmt.Printf("n is %d.- you've been shadowed", n)
}


