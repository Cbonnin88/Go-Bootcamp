package number

import (
	"fmt"
	"math/rand"
	"time"
)

func LuckyNumber() {

	source := rand.NewSource(time.Now().UnixNano())
	guess := rand.New(source).Intn(10)

	var try int

	for {
		fmt.Print("Please pick a number between 1 and 10: ")
		var num int
		scan, err := fmt.Scan(&num)
		if err != nil {
			panic(scan)
			return
		}

		try++

		switch true {
		case num > guess:
			fmt.Println("Too High")

		case num < guess:
			fmt.Println("Too Low")

		case num == guess:
			fmt.Printf("Great Guess, it only took you %d tries\n", try)
			return
		}

	}
}
