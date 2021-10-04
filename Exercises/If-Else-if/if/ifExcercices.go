package _if

import (
	"fmt"
)

func Retire() {
	fmt.Print("Enter the year you were born: ")
	var year int
	_, err := fmt.Scan(&year)
	if err != nil {
		panic("Please enter a valid year")
	}
	age := 2021 - year

	if age >= 65 {
		fmt.Println("You can Retire at 100% of your pension")
	} else if age < 65 && age >= 55{
		fmt.Println("You can retire but with 85% of you pension")

	} else if age < 55 && age >=18 {
		fmt.Println("You are ineligible for retirement ")
	} else {
		fmt.Println("Please enter a year starting from 2003")
	}
}

func Loan() {
	var (
		radius int
	)
	fmt.Print("Enter radius size: ")
	_, err := fmt.Scan(&radius)
	if err != nil {
		panic("Please enter a valid number")
	}

	if radius >= 200 {
		fmt.Println("it's a big sphere")

	} else if radius < 200 && radius >= 100 {
		fmt.Println("It's a normal sphere")
	} else {
		fmt.Println("I don't know")
	}
}

func Pets() {
	var pet string

	fmt.Print("What type of pet do you have: ")
	_, err := fmt.Scan(&pet)
	if err != nil {
		panic("Enter a valid word")
	}

	if pet == "Dog" || pet == "Cat" || pet == "bird" {
		fmt.Println("This is a typical pet")

	} else if pet == "Scorpion" || pet == "Spider" || pet == "Snake" {
		fmt.Println("This is dangerous pet !!!")

	} else if pet == "Lion" || pet == "Elephant" || pet == "Tiger" {
		fmt.Println("That isn't a pet")
	} else {
		fmt.Println("You own a pet that's not on my list")
	}
}

func Vowel() {
	var vowels string
	fmt.Print("Enter a letter from A to Z: ")
	_, err := fmt.Scan(&vowels)
	if err != nil {
		panic("Please enter a letter")
	}

	if vowels == "a" || vowels == "e" || vowels == "i" || vowels == "o" || vowels == "u" {
		fmt.Printf("%v is a vowel\n", vowels)
	} else if vowels == "y" {
		fmt.Printf("%v is sometimes a vowel\n", vowels)
	} else {
		fmt.Printf("%v is a consonant\n", vowels)
	}
}


