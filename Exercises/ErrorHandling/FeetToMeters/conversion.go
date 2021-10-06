package FeetToMeters

import (
	"fmt"
	"math"
)

func Convert() {
	fmt.Print("Enter your height in feet: ")
	var feet float64
	_, err := fmt.Scan(&feet)
	onlyPositive := math.Signbit(feet)

	if err != nil {
		fmt.Printf("Error: Please enter a number")
		return
	} else if onlyPositive {
		fmt.Printf("Error: %g is not a postive number",feet)
		return
	}

	meters := feet * 0.3048
	fmt.Printf("%g feet is %.2f meters.",feet,meters) // %g for floats
}
