package main

import (
	"Go-BootCamp/Examples/SwitchStatements/Default"
	"fmt"
)

func main() {

	// Quick example of switch statements:
	fmt.Print("Choose a Country from this list to find its' Capital:\n" +
		"1.France\n2.United States\n3.Japan\n4.South Korea\n\nYour choice: ")

	 var choice int

	_, err := fmt.Scan(&choice)
	if err != nil || choice > 4 {
		fmt.Println("Error: Please choose a country from the list")
		return
	}
	switch choice {
	case 1 :
		country := "France"
		fmt.Printf("The capital of %s is Paris", country)

	case 2 :
		country := "United States"
		fmt.Printf("The capital of %s is Washington D.C",country)

	case 3 :
		country := "Japan"
		fmt.Printf("The capital of %s is Tokyo",country)

	case 4 :
		country := "South Korea"
		fmt.Printf("The capital of %s is Seoul",country)
	}

	Default.Age()
	Default.MultipleCase()
	Default.Truth()
	Default.FallThrough()
	Default.ShortSwitch()

	}





