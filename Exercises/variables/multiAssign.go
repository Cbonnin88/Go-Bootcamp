package main

import "fmt"

func main() {
	var (
		lang  string
		version  int
		planet string
		isTrue bool
		temp   float64
	)

	lang = "go"
	version = 2
	fmt.Println(lang,"version",version)

	planet = "Air is good on mars"
	isTrue = true
	temp = 19.5
	fmt.Println(planet,"\nIt's",isTrue,"\nIt is",temp,"degrees.")

}
