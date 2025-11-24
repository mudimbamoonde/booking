package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets int
	remainingTickets = conferenceTickets
	greetUsers(conferenceName) // function call

	// fmt.Println("We have ",conferenceTickets, " tickets and ",remainingTickets," are still available")
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
	bookings := []string{}
	for {

		var (
			firstName   string
			lastName    string
			email       string
			userTickets int
		)
		// ask a user for their name
		fmt.Print("Enter Your First Name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter Your Last Name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter Your Email: ")
		fmt.Scan(&email)

		fmt.Print("How many Tickets do you want?  ")
		fmt.Scan(&userTickets)

		//var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//if !isValidEmail {
		//	fmt.Println("Invalid Email")
		//	continue
		//}
		//if !isValidName {
		//	fmt.Println("Invalid Name")
		//	continue
		//}
		if isValidTicketNumber && isValidEmail && isValidName {

			remainingTickets = remainingTickets - userTickets
			//bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Println("======================================================================================================================================")
			//fmt.Printf("The whole array %v \n", bookings)
			//fmt.Printf("The first value %v \n", bookings[0])
			//fmt.Printf("Array type %T \n", bookings)
			//fmt.Printf("Slice length %v \n", len(bookings))

			fmt.Printf("Thank you %v %v for  booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Println("======================================================================================================================================")

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}

			fmt.Printf("We have  %v remaining Tickets.\n", remainingTickets)
			//fmt.Printf("These are all our bookings: %v\n", bookings)
			fmt.Printf("The First names of  bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, all tickets have been sold out!")
				break
			}
		} else {

			if !isValidEmail {
				fmt.Println("Invalid Email")

			}
			if !isValidName {
				fmt.Println("Invalid Name")

			}
			if !isValidTicketNumber {
				fmt.Println("Invalid Ticket Number")
			}
		}

	}

}

func greetUsers(conferenceName string) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
}
