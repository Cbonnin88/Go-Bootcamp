package RawStrings

import (
	"fmt"
	"os" // Allow me to access the operating system functionalities
)

func Greeter() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path:", os.Args[0])
}
