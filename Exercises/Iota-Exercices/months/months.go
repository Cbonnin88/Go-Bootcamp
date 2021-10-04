package months

import (
	"fmt"
)

func Months() {
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
	fmt.Println("Months of the Year (reverse order):",Dec,Nov,Oct,Sept,Aug,Jul,June,May,Apr,Mar,Feb,Jan)
	fmt.Println("Months of the Year (natural order):",Jan,Feb,Mar,Apr,May,June,Jul,Aug,Sept,Oct,Nov,Dec)
}

func ShortMonths() {
	const (
		_ = iota
		Jan
		Feb
		Mar
	)
	fmt.Println("January, February, March:",Jan, Feb, Mar)
}

func Seasons() {
	const(
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)
	fmt.Println("Seasons:",Winter,Spring,Summer,Fall)
}
