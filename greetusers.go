package main

import "fmt"

func greetUsers() {
	fmt.Printf("welcome to my %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
	fmt.Println("---------------------------------------------------")
}