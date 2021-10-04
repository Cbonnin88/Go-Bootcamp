package main

import (
	"Go-BootCamp/DataStructures/BigO"
	"Go-BootCamp/DataStructures/SpaceComplexity"
	"fmt"
)

func main() {
	fmt.Println("Data Structures and Algorithms")
	BigO.One([]int{12,13,14,15})
	fmt.Println("")
	BigO.Two([]int{1,2,3,4}, []string{"Tim","Jane","Mark", "Alice"})
	fmt.Println("")
	BigO.Three([]int{1,2,3, 4})

	fmt.Println("Space Complexity:")
	SpaceComplexity.Greet([]string{"John", "Loki", "Alcina", "Ethan"})



}
