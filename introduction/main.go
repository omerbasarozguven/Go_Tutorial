package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"
	// OR
	conferenceName := "Go Conference"
	const conferanceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v boking application\n", conferenceName)
	fmt.Println("We have total of", conferanceTickets, "tickets and", remainingTickets, "are still avalible")
	fmt.Println("Get your tickets here to attend")
	//To print variable types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferanceTickets, remainingTickets, conferenceName)

	// var userName string
	// userName = "Tom"

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets to be booked: ")
	fmt.Scan(&userTickets)

	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)
}
