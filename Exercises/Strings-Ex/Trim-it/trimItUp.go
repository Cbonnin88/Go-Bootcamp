package Trim

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

<<<<<<< HEAD
func RawLiteral() {
=======
func Trim() {
>>>>>>> 19e7ad48599b17513039c3c773268e4fea2ba396
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
