package Trim

import "fmt"

func JSON() {
	json := `"\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n`

	fmt.Println(json)

}
