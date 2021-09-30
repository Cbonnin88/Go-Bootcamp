package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := `

	The weather looks good.
I should go and play.

	`
	trim := strings.TrimSpace(msg)
	fmt.Println(trim)

	name := "inanç           "
	fmt.Println(name)

	name = strings.TrimRight(name, " ")
	length := utf8.RuneCountInString(name)
	fmt.Println(length)


}
