package multiplicationTable

import "fmt"

func Multiply() {

	const max = 12

	fmt.Printf("%5s", "X") // will be five characters wide

	for i := 0; i <= 12; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for o := 0; o <= max; o ++ {
		fmt.Printf("%5d", o)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", o*j )
		}
		fmt.Println()
	}
	fmt.Println("")
}

func MultiplyAgain() {
	const max = 25

	fmt.Printf("%5s", "X") // will be five characters wide

	for i := 13; i <= 25; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for o := 13; o <= max; o ++ {
		fmt.Printf("%5d", o)

		for j := 13; j <= max; j++ {
			fmt.Printf("%5d", o*j )
		}
		fmt.Println()
	}

}

func MultiplyOnceMore() {
	fmt.Print("\n\nEnter a minimum table size: ")
	var min int
	scan, err := fmt.Scan(&min)
	if err != nil {
		fmt.Println(scan)
		return
	}

	fmt.Print("Enter a maximum table size: ")
	var max int
	scan, err = fmt.Scan(&max)
	if err != nil {
		fmt.Println(scan)
		return
	}

	fmt.Println("You multiplication table:")
	fmt.Printf("%5s", "X")

	for i := min; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for o := min; o <= max; o ++ {
		fmt.Printf("%5d", o)

		for j := min; j <= max; j++ {
			fmt.Printf("%5d", o*j )
		}
		fmt.Println()
	}
	fmt.Println("")




}