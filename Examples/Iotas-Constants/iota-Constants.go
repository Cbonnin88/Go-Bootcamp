package main

import "fmt"

func main() {
	const (
		monday  = iota + 1 // starts at zero
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday,thursday,friday,saturday,sunday)

	const (
		EST = -(5 + iota)
		_ // Blank Identifier
		MST
		PST
	)
	fmt.Println(EST,MST,PST)

}
