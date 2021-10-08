package breaks

import "fmt"

func Break() {
	var (
		sum int
		i = 1
	)

	for {
		if i > 10 {
			break // Here we break the loop
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}
