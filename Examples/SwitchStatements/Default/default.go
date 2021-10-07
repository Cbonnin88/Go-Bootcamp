package Default

import "fmt"

func Age() {
	fmt.Print("\nEnter your age: ")
	var age int

	_, err := fmt.Scan(&age)
	if err != nil || age < 0 {
		fmt.Println("Error: Please enter a valid number")
		return
	}

	switch age {
	case 18 :
		fmt.Println("You are able to drive with no restrictions")

	case 16, 17 :
		fmt.Println("You are able to drive, but only until midnight")

	case 15 :
		fmt.Println("You can apply for your learners permet")

	default:
		fmt.Println("You cannot drive")
	}

}

func MultipleCase() {
	fmt.Print("Enter a name of a city: ")
	var city string
	_, err := fmt.Scan(&city)
	if err != nil {
		fmt.Println("Please enter a city name")
		return
	}

	switch city {
	case "Paris", "Lyon":
		fmt.Printf("%s is located in France",city)

	case "Tokyo", "Kyoto":
		fmt.Printf("%s is located in Japan", city)

	case "Philadelphia", "New York":
		fmt.Printf("%s is located in the United States",city)

	case "London", "Maidenhead":
		fmt.Printf("%s is located in the United Kingdom",city)

	default:
		fmt.Println("Where ?")
	}
}

func Truth() {
	fmt.Print("\nEnter a number: ")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error: Enter a number")
		return
	}

	switch true {
	case number > 0 :
		fmt.Printf("%d is a positive number", number)

	case number < 0:
		fmt.Printf("%d is a negative number", number)

	default:
		fmt.Println("zero")
	}
}

func FallThrough() {
	fmt.Print("\nEnter a number: ")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error: Enter a number")
		return
	}

	switch true {
	case num >= 1000:
		fmt.Println("A huge number")
		fallthrough

	case num >= 100 && num < 1000:
		fmt.Println("A big number")
		fallthrough

	case num > 0 && num < 100:
		fmt.Println("a positive number")
		fallthrough

	default:
		fmt.Println("A number")
	}
}

func ShortSwitch() {
	switch i := 12; {
	case i > 0 :
		fmt.Println("positive")

	case i < 0 :
		fmt.Println("Negative")

	default:
		fmt.Println("Zero")
	}
}