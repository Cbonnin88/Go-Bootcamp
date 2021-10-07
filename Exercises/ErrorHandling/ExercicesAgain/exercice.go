package ExercicesAgain

import (
	"fmt"
	"strings"
	"time"
)

func MovieRating() {
	fmt.Print("\nHow old are you: ")
	var age int
	_, err := fmt.Scan(&age)
	if err != nil || age < 0 {
		fmt.Println("Please enter a valid age",)
		return
	}

	if age >= 17 {
		fmt.Println("You age able to see a R-rated Movie")
	} else if age >= 13 && age <= 17 {
		fmt.Println("You are able to see a PG-13 movie")
	} else if age < 13 {
		fmt.Println("You are ale to see a PG-Movie")
	}
}

func OddOrEven() {
	fmt.Print("Enter a number: ")
	var number int
	_, err := fmt.Scan(&number)

	if err != nil || number < 0 || number != number {
		fmt.Println("Error: Please enter a valid number")
		return 
	}

	if number%8 == 0 {
		fmt.Printf("%d is an even number and is divisible by 4", number)
	} else if number%2 == 0{
		fmt.Printf("%d is an even number", number)
	} else {
		fmt.Printf("%d is an odd number", number)
	}
}

func LeapYear() {
	fmt.Print("\nEnter a valid year from 2000 onwards: ")
	var year int
	_, err := fmt.Scan(&year)
	if err != nil || year < 2000 {
		fmt.Println("Enter a year starting from 2000")
		return
	}

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

func DaysOfTheMonth() {
	fmt.Print("Enter a month: ")
	var month string
	_, err := fmt.Scan(&month)
	if err != nil {
		fmt.Println("Error: Please enter a month")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)
	days := 28


	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30

	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31

	} else if m == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
