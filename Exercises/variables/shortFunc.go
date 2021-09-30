package main

import "fmt"

func main() {
	a,_ := multi()
	fmt.Println(a)

	_,b := multi2()
	fmt.Println(b)
}


func multi() (int, int)  {
	return 5,4
}

func multi2() (string, string) {
	return "Hello", "There"
}
