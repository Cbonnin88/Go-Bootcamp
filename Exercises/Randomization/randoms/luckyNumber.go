package randoms

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessingGame(){
	rand.Seed(time.Now().UnixNano())
	guess := rand.Intn(10)
	var try int

	for {

		fmt.Print("Pick a number between 1 and 10: ")
		var num int
		scan, err := fmt.Scan(&num)
		if err != nil {
			panic(scan)
			return
		}
		try++

		switch true {
		case num > guess:
			fmt.Println("Too High, try again")

		case num < guess:
			fmt.Println("Too low, try again")

		case num == guess && try == 1:
			fmt.Println("Awesome, you got it on the first try !!!")
			return

		case num == guess:
			fmt.Printf("Awesome!!!, you got it and it only took you %d tries\n", try)
			return
		}
	}

}


