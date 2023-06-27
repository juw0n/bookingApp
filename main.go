package main

import (
	"fmt"
	"strings"
)

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

	var bookings []string

	for {
		// Getting users details
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Printf("How many ticket are you ordering: ")
		fmt.Scan(&userTickets)

		validName := len(firstName) >= 2 && len(lastName) >= 2
		validEmail := strings.Contains(email, "@")
		validTicket := userTickets > 0 && remainingTickets >= userTickets

		// some Logics
		if validName && validEmail && validTicket {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for ordering %v tickets. You will get a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tickets left for %v.\n", remainingTickets, confName)

			if remainingTickets == 0 {
				fmt.Println("Our conference tickets are sold our. you can view on our website")
				break
			}

		} else {
			if !validName {
				fmt.Println("Your firstname or lastname is invalid")
			}
			if !validEmail {
				fmt.Println("Invalid email address")
			}
			if !validTicket {
				fmt.Println("You ticket number is invalid")
			}
			continue
		}

		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all the bookings: %v\n", firstNames)
	}

}
