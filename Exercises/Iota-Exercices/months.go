package main

import (
	"fmt"
)

func main() {
	const (
		Dec = 12 - iota
		Nov
		Oct
		Sept
		Aug
		Jul
		June
		May
		Apr
		Mar
		Feb
		Jan
	)
	fmt.Println(Dec,Nov,Oct,Sept,Aug,Jul,June,May,Apr,Mar,Feb,Jan)
	fmt.Println(Jan,Feb,Mar,Apr,May,June,Jul,Aug,Sept,Oct,Nov,Dec)

}
