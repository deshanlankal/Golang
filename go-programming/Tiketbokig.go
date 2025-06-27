package main

import (
	"fmt"
	"strings"
)

func main() {
	eventName := "Concert"
	const eventTikets int = 50
	var remainingEventTikets uint = 50
	// var bookings = [50]string{"john doe","jane smith","alice johnson", "bob brown", "charlie green", "dave white", "eve black", "frank gray", "george blue", "harry red",}
	// var bookings [50]string 
	var bookings []string

	fmt.Printf("Welcome to the %v event!\n", eventName)
	fmt.Printf("We got total of %d tickets available.\n", eventTikets)
	fmt.Printf("Get your event Pasese today before they sold out")


	for{

		var firstName string
		var lastName string
		var email string
		var userTikets uint

		

		fmt.Print("Enter your first name: ")
		fmt.Scanln(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scanln(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scanln(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scanln(&userTikets)

		remainingEventTikets =  remainingEventTikets - userTikets
		// bookings[0] = firstName + " " + lastName
		/// Use of Dynamic list, set the "var bookings [50]string" the size to "[ ]" this
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Print("Thank you for your purchase, ", firstName, " ", lastName, "!\n")
		fmt.Printf("A confirmation email has been sent to %s.\n", email)
		fmt.Printf("You have successfully purchased %d tickets for the %v event.\n", userTikets, eventName)
		fmt.Printf("Enjoy the event and have a great time!\n")
	///////////// Array Info ///////////////////
		fmt.Printf("The total num of bookings so far:%s\n", bookings)
		fmt.Printf("the first of the slice%v\n", bookings[0])
		fmt.Printf("The bookig slice type%T\n", bookings)
		fmt.Printf("Booking slice length:%v\n", len(bookings))
	////////////////////////////////////////////////////////////

		firstNames := []string{} 
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First names of the bookings: %v\n", firstNames)
		
		fmt.Printf("Remaining tickets for the %v event: %d\n", eventName, remainingEventTikets)
	}
	
}