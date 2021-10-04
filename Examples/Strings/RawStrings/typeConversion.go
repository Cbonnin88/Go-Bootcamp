package RawStrings

import "fmt"

func Conversion() {
	speed := 200
	force := 2.5

	speed = int(float64(speed) * force)
	fmt.Println(speed)


}
