package loopingSlices

import (
	"fmt"
	"strings"
)

func Sentence() {
	words := strings.Fields("Rest while you can, because I will hunt you and I will break you !!!")

	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}
}

// -2d means printing the left align
// strings.Fields() splits this string from teh white-space points