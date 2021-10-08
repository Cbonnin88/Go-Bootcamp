package continuing

import "fmt"

func Continuing() {

	var (
		i = 1
		sum int
	)

	for {
		if i > 20 {
			break
		}

		if i % 2 != 0 {
			i++
			continue
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
}