package main

import "fmt"

func main() {
	var firstName string
	var lastName string

	fmt.Printf("Please input your first name :")
	fmt.Scanln(&firstName)

	fmt.Printf("Please input your last name :")
	fmt.Scanln(&lastName)

	fmt.Println("Your full name " + firstName + " " + lastName)

}
