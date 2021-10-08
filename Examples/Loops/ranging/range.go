package ranging

import (
	"fmt"
	"os"
)

func Range() {
	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}
}

