package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]

	s := msg + strings.Repeat("!", len(msg))
	fmt.Println(s)

}
