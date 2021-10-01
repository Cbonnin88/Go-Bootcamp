package main

import "fmt"

func main() {

	// printf = %q
	var brand = "Molotov"
	fmt.Printf("%q\n", brand)

	// %s, without double quotes
	var car = "Audi"
	fmt.Printf("%s\n", car)

	// %d, interger value
	success := 545
	failure := 400
	totalRateRatio := success / failure
	fmt.Printf("Success: %d\nFailures: %d\nTotal Success Ratio: %d",success,failure,totalRateRatio)
}
