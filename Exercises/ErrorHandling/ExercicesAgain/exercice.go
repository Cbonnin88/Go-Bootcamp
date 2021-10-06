package ExercicesAgain

import (
	"fmt"
	"math"
)

func MovieRating() {
	fmt.Print("\nHow old are you: ")
	var age int
	age, err := fmt.Scan(&age)
	onlyPositive := math.Signbit(float64(age))
	if err != nil {
		fmt.Println("Error: Please enter a valid number")
		return
	} else if onlyPositive{
		fmt.Println("Error: Please enter a positive number")
		return
	}

	if age > 17 {
		fmt.Println("You are able to see R-Rated movies")
	} else if age >= 13 && age <= 17 {
		fmt.Println("You are able to see PG-13 movies")

	} else if age < 0 {
		fmt.Println("You are able to see a PG-Rated movie")
	}

}
