package SpaceComplexity

import "fmt"

func Greet(names []string) {
	// O(1) space
	for n := 0; n < len(names); n++ {
		fmt.Println("Hello There " + names[n])
	}
}

// The more items that we have in our array, the more space that it will take
