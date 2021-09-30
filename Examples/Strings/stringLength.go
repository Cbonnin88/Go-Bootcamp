package main

import (
	"fmt"
	"unicode/utf8" // helps with utf8 encoded values
)

func main() {
	firstName := "Christopher"
	middleName := "Brian"
	lastName := "Bonnin"
	birthName := "Carrington"
	fullName := firstName + middleName + lastName
	withBirth := firstName + middleName + lastName + birthName
	foreignName := "ênaôçà"

	fmt.Println(firstName, "has", len(firstName), "letters.")
	fmt.Println(middleName, "has", len(middleName), "letters.")
	fmt.Println(lastName, "has", len(lastName), "letters.")
	fmt.Println(birthName, "has", len(birthName), "letters.")
	fmt.Println("My full name as a total of",len(fullName))
	fmt.Println("With my birth name included", len(withBirth))
	fmt.Println("This foreign name has", utf8.RuneCountInString(foreignName))

	// utf8 packages RuneCountingString function finds the number of characters in a string
	// utf8 helps with non english values

}
