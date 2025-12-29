package main

import "strings"

func validticket (FirstName string, LastName string, email string, city string, userTickets uint, remainingTickets uint)(bool, bool, bool, bool) {

	validName := len(FirstName) >=2 && len(LastName) >=2
	validEmail := strings.Contains(email, "@")
	validTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	validCity := strings.Contains(city,"Singapore") || strings.Contains(city,"New York") || strings.Contains(city,"London") || strings.Contains(city,"Berlin")

	return validName, validEmail, validTicketNumber, validCity
}