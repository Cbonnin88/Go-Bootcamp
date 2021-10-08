package main

import (
	"Go-BootCamp/Examples/Loops/breaks"
	"Go-BootCamp/Examples/Loops/continuing"
	"Go-BootCamp/Examples/Loops/loopingSlices"
	"Go-BootCamp/Examples/Loops/ranging"
	"fmt"
)

func main() {

	// Simple Loop statement
	var sum int

	for i := 1; i <= 5; i++ {
		sum += i
	}

	fmt.Println(sum)
	breaks.Break()
	continuing.Continuing()
	loopingSlices.Sentence()
	ranging.Range()
}
