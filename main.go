package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// Slice
	var bookings []string

	// Greet the users to our ticket application.
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
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

		//Field Validations
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// Check if user is trying to book more tickets than available.
		if isValidName && isValidEmail && isValidTicketNumber {
			// Subtract the tickets booked.
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
			fmt.Printf("%d tickets now remain for the %s\n", remainingTickets, conferenceName)

			// Call function to print the users that registered for the conference; first names only.
			printFirstNames(bookings, conferenceName)

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
func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

// Function prints registered conference user's first names only.
func printFirstNames(bookings []string, conferenceName string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These first names of the bookings for the %s: %s\n", conferenceName, firstNames)
}
