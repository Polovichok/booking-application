package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainigTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 &&  userTickets <= remainigTickets
	return isValidName, isValidEmail, isValidTickets
}