package SwitchStatements

import (
	"fmt"
	"strings"
	"time"
)

func TimeOfTheDay() {
	switch h := time.Now().Hour(); {
	case h == 7:
		fmt.Println("Time to wake up")

	case h == 8:
		fmt.Println("Time to leave")

	case h == 9:
		fmt.Println("Start the work day")

	case h == 12:
		fmt.Println("Lunch Break")

	case h == 14:
		fmt.Println("Lunch Break over")

	case h == 16:
		fmt.Println("Afternoon Meeting")

	case h == 18:
		fmt.Println("Work day finished")

	case h == 20:
		fmt.Println("Finally home")


	case h == 23:
		fmt.Println("Bedtime")
	}
}

func RichterScale() {
	fmt.Print("Enter the earthquake scale: ")
	var eq string

	_, err := fmt.Scan(&eq)
	if err != nil {
		fmt.Println("Error: Enter a word")
		return
	}

	var richter string

	switch eq {
	case "Micro":
		richter = "Less than 2.0"

	case "Very Minor":
		richter = "2 - 2.9"

	case "Minor":
		richter = "2 - 2.9"

	case "Light":
		richter = "4 - 4.9"

	case "Moderate":
		richter = "5 - 5.9"

	case "Strong":
		richter = "6 - 6.9"

	case "Major":
		richter = "7 - 7.9"

	case "Great":
		richter = "8 - 9.9"


	case "Massive":
		richter = "10+"

	default:
		richter = "Unknown"
	}
	fmt.Printf("%s richter scale is %s\n", eq, richter)
}

func Password() {
	fmt.Print("Enter you User ID: ")
	var id string
	scan, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println(scan)
		return 
	}
	
	fmt.Print("Enter your password: ")
	var pwd string
	scan, err = fmt.Scan(&pwd)
	if err != nil {
		fmt.Println(scan)
		return 
	}

		name := "Christopher Bonnin"



	switch true {
	case id == "cbonnin" && pwd == "molotov":
		fmt.Printf("Access granted for %q",name)

	case id != "cbonnin":
		fmt.Println("Invalid user name")

	case pwd != "molotov":
		fmt.Println("Invalid Password")
	}
}

func StringManipulator() {
	fmt.Print("\nEnter, lower, upper or title to see changes: ")
	var change string
	scan, err := fmt.Scan(&change)
	if err != nil {
		fmt.Println("Error: Unknown Command",scan)
		return
	}

	fmt.Print("Please type the word you would like to see changed: ")
	var word string
	scan, err = fmt.Scan(&word)
	if err != nil {
		fmt.Println(scan)
		return
	}

	switch  {
	case change == "upper":
		fmt.Println(strings.ToUpper(word))

	case change == "lower":
		fmt.Println(strings.ToLower(word))

	case change == "title":
		fmt.Println(strings.ToTitle(word))
	}
}




