package printFExercice

import (
	"fmt"
)

func Age()  {
	myAge := 30
	fmt.Printf("You should change %d to your age, of course\n", myAge)
}

func FullName() {
	first := "Christopher"
	last := "Bonnin"
	fmt.Printf("My first and last name is %s %s\n", first, last)
	fmt.Printf("'first' is of type %T\n", first)
}

func Claims() {
	fmt.Printf("These are %t claims.\n", false)
}

func Temp() {
	t := 95.676765753
	fmt.Printf("The temperature outside is %.1f in Philadelphia\n", t)
}

func HelloWorld() {
	word := "Rest while you can, because I will hunt you, and I will break you !!!!"
	fmt.Printf("%s\n",word)
}

func Type() {
	num := 3
	num2 := 3.14
	word := "Hello"

	fmt.Printf("'num' is of type %T\n", num)
	fmt.Printf("'num2 is of type %T\n", num2)
	fmt.Printf("'word' is of type %T\n", word)
	fmt.Printf("'false' is of type %T\n", false)
}

func Arguments() {
	first, last := "Christopher", "Bonnin"
	msg := "Your first and last name is %s %s\n"
	fmt.Printf(msg, first,last)
}
