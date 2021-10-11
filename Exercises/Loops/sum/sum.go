package sum

import (
	"fmt"
	"math/rand"
	"time"
)

func Sum() {
	var sum int

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func VerboseSum() {
	const min, max = 1, 10
	var sum int

	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

func CommandLine() {
	fmt.Print("Enter a minimum value: ")
	var min int
	scan, err := fmt.Scan(&min)
	if err != nil {
		panic(scan)
		return
	}

	fmt.Print("Enter a maximum value: ")
	var max int
	scan, err = fmt.Scan(&max)
	if err != nil {
		panic(scan)
		return
	}

	var sum int

	for o := min; o <= max; o++ {
		sum += o

		fmt.Print(o)
		if o < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n",sum)
}

func Even() {
	fmt.Print("Enter a minimum value: ")
	var min int
	scan, err := fmt.Scan(&min)
	if err != nil {
		panic(scan)
		return
	}

	fmt.Print("Enter a maximum value: ")
	var max int
	scan, err = fmt.Scan(&max)
	if err != nil {
		panic(scan)
		return
	}

	var sum int

	for j := min; j <= max; j++ {
		if j%2 != 0 {
			continue
		}
		sum += j

		fmt.Print(j)
		if j < max - 1 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n",sum)
}

func BreakUp() {

	fmt.Print("Enter a minimum value: ")
	var min int
	scan, err := fmt.Scan(&min)
	if err != nil {
		panic(scan)
		return
	}

	fmt.Print("Enter a maximum value: ")
	var max int
	scan, err = fmt.Scan(&max)
	if err != nil {
		panic(scan)
		return
	}

	var(
		sum int
		j int
	)

	for  {
		if j > max {
			break
		} else if j%2 != 0 {
			j++
			continue
		}

		fmt.Print(j)
		if j < max - 1 {
			fmt.Print(" + ")
		}
		sum += j
		j++
	}
	fmt.Printf(" = %d\n",sum)
}

func InfiniteKill() {
	for {
		var c string

		switch rand.Intn(5){
		case 0:
			c = "Hi"
		case 1:
			c = "There"
		case 2:
			c = "You"
		case 3:
			c = "My"
		case 4:
			c = "Friend"
		}
		fmt.Printf("\r%s Please Wait. Processing....",c)
		time.Sleep(time.Millisecond * 250)
		break

	}
}

