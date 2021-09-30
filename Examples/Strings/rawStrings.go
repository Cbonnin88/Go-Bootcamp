package main

import "fmt"

func main() {
	regString := "Rest while you can, because I will hunt you, and I will break you!!!"
	rawString := `Rest while you can, because I will hunt you, and I will break you!!!`
	fmt.Println(regString)
	fmt.Println(rawString)

	h := `
 <html>
     <body>"Hello There"</body>
 <html>`
	fmt.Println(h)

	file := `c:\usr\dir\file`
	fmt.Println(file)

}
