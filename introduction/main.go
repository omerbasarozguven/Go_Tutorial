package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var dummy = "global variable"

func main() {
	// var conferenceName string = "Go Conference"
	// OR
	conferenceName := "Go Conference"
	const conferanceTickets int = 50
	var remainingTickets int = 50
	var bookingsSlice []string

	fmt.Println(dummy)
	fmt.Println(helper.HelperVar)

	fmt.Printf("Welcome to %v boking application\n", conferenceName)
	fmt.Println("We have total of", conferanceTickets, "tickets and", remainingTickets, "are still avalible")
	// fmt.Println("Get your tickets here to attend")
	helper.PrintFnc("Get your tickets here to attend")

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

		// if userTickets > remainingTickets {
		// 	fmt.Println("We only have", remainingTickets, "remaining, so you can't book", userTickets, "tickets!")
		// 	fmt.Println("Please try again")
		// 	continue
		// }

		for userTickets > remainingTickets {
			fmt.Println("We only have", remainingTickets, "remaining, so you can't book", userTickets, "tickets!")
			fmt.Print("Enter number of tickets to be booked: ")
			fmt.Scan(&userTickets)
		}

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

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out come back next year")
			break
		}
	}

	// MAPS
	var firstMap = make(map[string]string)
	firstMap["firstName"] = "omer"
	firstMap["lastName"] = "ozguven"
	firstMap["email"] = "omer@gmail.com"
	fmt.Println(firstMap)
	var mapSlice = make([]map[string]string, 0)
	mapSlice = append(mapSlice, firstMap)
	fmt.Println(mapSlice)
}
