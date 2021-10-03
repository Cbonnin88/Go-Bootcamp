package printf

import "fmt"

func PrintF() {
	// printf = %q
	var brand = "Molotov"
	fmt.Printf("%q\n", brand)

	// %s, without double quotes
	var car = "Audi"
	fmt.Printf("%s\n", car)

	// %d, integer value
	success := 25
	failure := 5
	totalRateRatio := success / failure
	fmt.Printf("Success: %d\nFailures: %d\nTotal Success Ratio: %d\n",success,failure,totalRateRatio)
}

func Escape() {
	// \n = new line
	fmt.Println("hi\nhi")
}
