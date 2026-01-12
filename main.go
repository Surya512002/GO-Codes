package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 1000

var remainingTickets uint = 1000
var bookings = make([]UserData, 0)

type UserData struct {
	FirstName   string
	LastName    string
	email       string
	city        string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//Greeting the user
	greetUsers()

	FirstName, LastName, email, city, userTickets := getUserInput()

	validName, validEmail, validTicketNumber, validCity := validticket(FirstName, LastName, email, city, userTickets, remainingTickets)

	if validName && validEmail && validTicketNumber && validCity {

		remainingTickets, bookings = bookingtickets(FirstName, LastName, email, city, userTickets, remainingTickets, conferenceName, bookings)

		//create a map of user data

		PrintFirstNames(bookings)

		wg.Add(1)
		go sendingticket(userTickets, FirstName, LastName, email)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			//break

		}
	} else {
		if !validName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !validEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !validTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
		if !validCity {
			fmt.Println("City you entered is not in the list of event locations")
		}
		wg.Wait()
	}
}

func getUserInput() (string, string, string, string, uint) {
	var FirstName string
	var LastName string
	var email string
	var city string
	var userTickets uint

	//asking for user input
	fmt.Println("Enter your First name: ")
	fmt.Scan(&FirstName)
	fmt.Println("Enter your Last name: ")
	fmt.Scan(&LastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("This Event is happening in Singapore,New York,London,Berlin.Choose Your City: ")
	fmt.Scan(&city)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return FirstName, LastName, email, city, userTickets
}

func PrintFirstNames(bookings []UserData) (FirstNames []string) {

	for _, booking := range bookings {
		FirstNames = append(FirstNames, booking.FirstName)

	}
	return FirstNames

}
func bookingtickets(FirstName string, LastName string, email string, city string, userTickets uint, remainingTickets uint, conferenceName string, bookings []UserData) (uint, []UserData) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		FirstName:   FirstName,
		LastName:    LastName,
		email:       email,
		city:        city,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FirstName, LastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets, bookings

}

func sendingticket(userTickets uint, FirstName string, LastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, FirstName, LastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}
