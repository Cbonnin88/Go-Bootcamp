package Yell_it_back

import (
	"fmt"
	"os"
	"strings"
)

func Yell() {

	msg := os.Args[1]
	length := len(msg)
	s := strings.Repeat("!", length) + msg + strings.Repeat("!", length)
	s = strings.ToUpper(s)
	fmt.Println(s)

}
