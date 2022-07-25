package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conference"
	// OR
	conferenceName := "Go Conference"
	const conferanceTickets int = 50
	var remainingTickets int = 50
	var bookingsSlice []string

	fmt.Printf("Welcome to %v boking application\n", conferenceName)
	fmt.Println("We have total of", conferanceTickets, "tickets and", remainingTickets, "are still avalible")
	fmt.Println("Get your tickets here to attend")
	//To print variable types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferanceTickets, remainingTickets, conferenceName)

	// var userName string
	// userName = "Tom"

	for {
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

		remainingTickets = remainingTickets - userTickets

		fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)
		fmt.Println(remainingTickets, "tickets remaining for", conferenceName)

		// Array
		// var bookingsArr [50]string
		// bookingsArr[0] = firstName + " " + lastName
		// fmt.Println("The whole Array:", bookingsArr)
		// fmt.Println("The first value:", bookingsArr[0])
		// fmt.Printf("Array type: %T\n", bookingsArr)
		// fmt.Println("Array length:", len(bookingsArr))

		// Slice
		bookingsSlice = append(bookingsSlice, firstName+" "+lastName)
		// fmt.Println("The whole Slice:", bookingsSlice)
		// fmt.Println("The first value:", bookingsSlice[0])
		// fmt.Printf("Slice type: %T\n", bookingsSlice)
		// fmt.Println("Slice length:", len(bookingsSlice))

		// for each
		firstNames := []string{}
		for _, booking := range bookingsSlice {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Println("The first names of bookings are:", firstNames)
	}

}
