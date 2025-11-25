package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Pakage level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = conferenceTickets

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// function call
	greetUsers()
	// fmt.Println("We have ",conferenceTickets, " tickets and ",remainingTickets," are still available")1
	//for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidTicketNumber, isValidEmail, isValidName := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	if isValidTicketNumber && isValidEmail && isValidName {
		bookTicket(userTickets, firstName, lastName, email) // Book tickets

		// the go routine makes the application concurrent
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email) //Send emails,
		// call function print firstName
		firstNames := getFirstNames()
		fmt.Printf("The First names of  bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, all tickets have been sold out!")
			//break
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
	wg.Wait() // wait for all goroutines to finish
}

//}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//names := strings.Fields(booking)

		//firstNames = append(firstNames, names[0])
		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	//fmt.Printf("These are all our bookings: %v\n", bookings)
	//fmt.Printf("The First names of  bookings: %v\n", firstNames)

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var (
		firstName   string
		lastName    string
		email       string
		userTickets uint
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//Create a map for a user
	//var userData = make(map[string]string)
	//Struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTikects"] = strconv.FormatUint(uint64(userTickets), 10)

	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)

	fmt.Printf("Lis of bookings: %v \n", bookings)

	fmt.Println("======================================================================================================================================")
	//fmt.Printf("The whole array %v \n", bookings)
	//fmt.Printf("The first value %v \n", bookings[0])
	//fmt.Printf("Array type %T \n", bookings)
	//fmt.Printf("Slice length %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for  booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Println("======================================================================================================================================")
	fmt.Printf("We have  %v remaining Tickets.\n", remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("**************************************")
	fmt.Printf("Sending  %v to email address %v \n", ticket, email)
	fmt.Println("**************************************")
	wg.Done()
}
