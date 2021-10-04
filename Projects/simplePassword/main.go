package main

import "fmt"

func main() {
	var userName string
	var pwd int


	fmt.Print("Enter your User Name: ")
	user, err := fmt.Scan(&userName)
	if err != nil {
		panic(user)
	}

	fmt.Print("Enter your password: ")
	pass, err := fmt.Scan(&pwd)
	if err != nil {
		panic(pass)
	}

	if len(userName) != 5 && pwd != 5 {
		fmt.Println("Please enter a user name and password with at least five characters")

	} else {
		fmt.Println("Login Successful")
	}


}
