package main

import "fmt"

func main() {

	// printf = %q
	var brand = "Molotov"
	fmt.Printf("%q\n", brand)

	// %s, without double quotes - string values
	var car = "Audi"
	fmt.Printf("%s\n", car)

	// %d, interger value
	success := 545
	failure := 400
	totalRateRatio := success / failure
	fmt.Printf("Success: %d\nFailures: %d\nTotal Success Ratio: %d\n",success,failure,totalRateRatio)

	// %T, type value
	var speed int
	fmt.Printf("%T\n", speed)

	// %v, print any value
	var (
		planet = "Venus"
		distance = 261
		orbit = 224.701
		hasLife = false
	)
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d\n", distance)
	fmt.Printf("Orbit: %.1f\n", orbit) // %f is for a float value
	fmt.Printf("Does %s have life: %t\n",planet,hasLife)

	// Argument indexing : You can refer to past values by position
	fmt.Printf("%v is %v km away", planet, distance)

	// \n = new line
	fmt.Println("hi\nhi")
	fmt.Println("hi\n\"hi\"")



}
