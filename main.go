package main

import (
	"fmt"
	"strings"
)

// Package level variables so all functions and main can access them.
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// Slice
var bookings = []string{}

func main() {

	// Greet the users to our ticket application.
	greetUsers()

	for {
		// Retrieve the user input.
		firstName, lastName, email, userTickets := getUserInput()

		// Call function for validation.
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// Check if user is trying to book more tickets than available.
		if isValidName && isValidEmail && isValidTicketNumber {

			// Book the tickets to the conference and calculate the remaining tickets.
			bookTicket(userTickets, firstName, lastName, email)

			// Call function to print the users that registered for the conference; first names only.
			//printFirstNames(bookings)
			firstNames := getFirstNames()
			fmt.Printf("These first names of the bookings for the %s: %s\n", conferenceName, firstNames)

			if remainingTickets == 0 {
				// end the program.
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
		} else {
			// Show what was invalid
			if !isValidName {
				fmt.Println("First name or last name you entered was too short.")
			}
			if !isValidEmail {
				fmt.Println("Email you entered does not contain the @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered was invalid.")
			}
		}
	}
} // End main

// Function greets the users.
func greetUsers() {
	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

// Function prints registered conference user's first names only.
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
	//fmt.Printf("These first names of the bookings for the %s: %s\n", conferenceName, firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// Subtract the tickets booked.
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets now remain for the %s\n", remainingTickets, conferenceName)
}
