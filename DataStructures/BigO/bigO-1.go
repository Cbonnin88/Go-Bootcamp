package BigO

import "fmt"

func One(numbers []int) {
	// O(1) - constant type
	fmt.Println("O(1): ",numbers[0])
}

func Two(numbers []int, names []string) {
	// O(n) - n represents the size of the input

	for i := 0; i < len(numbers); i++ { // O(n)
		fmt.Println("O(n): ", numbers[i])

		for i := 0; i < len(names); i++ { // O(m) - distinguish between the two names
			fmt.Println("O(m): ", names[i])
		}
	}
}

func Three(numbers []int) {
	// O(n^3)
	for i := 0; i < len(numbers); i++ { // O(n)
		fmt.Println(numbers)
	}
	for n := 0; n < len(numbers); n++ { // O(n)
		for nu := 0; nu < len(numbers); nu++ { // O(n)
			for num := 0; num < len(numbers); num++ { // O(n)
				fmt.Println(n, ", ", nu, ", ", num)
			}
		}
	}
}

