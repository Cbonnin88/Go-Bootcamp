package user

import "fmt"

func Validate() {
	fmt.Print("Enter your user name: ")
	var useN string
	_, err := fmt.Scan(&useN)
	if err != nil {
		fmt.Println("Invalid format")
	}

	fmt.Print("Enter your password: ")
	var pwd string
	_, err = fmt.Scan(&pwd)
	if err != nil {
		panic("Invalid format")
	}

	if useN != "cbonnin" {
		fmt.Print("Invalid User Name")
	} else if pwd != "molotov" {
		fmt.Println("Invalid Password")
	} else {
		fmt.Println("Welcome Christopher Bonnin")
	}
}

func MultiUser() {
	fmt.Print("User one's user name:")
	var user1 string
	_, err := fmt.Scan(&user1)
	if err != nil {
		fmt.Println("Invalid format")
	}

	fmt.Print("User two's user name: ")
	var user2 string
	_, err = fmt.Scan(&user2)
	if err != nil {
		fmt.Println("Invalid format")
	}

	fmt.Print("User one's password: ")
	var pwd1 string
	_, err = fmt.Scan(&pwd1)
	if err != nil {
		panic("Invalid format")
	}

	fmt.Print("User two's password: ")
	var pwd2 string
	_, err = fmt.Scan(&pwd2)
	if err != nil {
		panic("Invalid format")
	}

	if user1 != "cbonnin" || user2 != "jcarrington" {
		fmt.Println("Invalid user name")
	} else if pwd1 != "molotov" || pwd2 != "social" {
		fmt.Println("Invalid Password")
	} else {
		fmt.Println("Welcome Mr. Bonnin and Mr. Carrington")
	}
}
