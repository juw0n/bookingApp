package main

import "fmt"

func main() {
	confName := "Go Conference"
	const confTickets int = 50
	remainingTickets := 50

	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining\n", confTickets, remainingTickets)
	fmt.Println("Get your access ticket here")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	var booking [50]string

	// Getting users details
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Printf("How many ticket are you ordering: ")
	fmt.Scan(&userTickets)

	// some Logics
	remainingTickets = remainingTickets - userTickets
	booking[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for ordering %v tickets. You will get a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("The remaining available ticket for this confrerence is %v.\n", remainingTickets)
}
