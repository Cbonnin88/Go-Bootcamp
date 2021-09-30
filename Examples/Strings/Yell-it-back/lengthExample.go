package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	msg := os.Args[1]
	length := len(msg)
	s := strings.Repeat("!", length) + msg + strings.Repeat("!", length)
	s = strings.ToUpper(s)
	fmt.Println(s)

}
