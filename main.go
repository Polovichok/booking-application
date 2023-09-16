package main

import "fmt"

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainigTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}


func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUsersInput()

		isValidName, isValidEmail,isValidTickets := validateUserInput(firstName, lastName, email, userTickets, remainigTickets)


		if  isValidName && isValidEmail && isValidTickets{
 
			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

	    	if remainigTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName{
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email adress you entered does not contain @ sign.")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable\n", conferenceTickets, remainigTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames 
}

func getUsersInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your  email adress: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets) 

	return firstName, lastName, email, userTickets
}

func bookTicket( userTickets uint, firstName string, lastName string, email string) {
	remainigTickets = remainigTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email )
	fmt.Printf("%v tickets remainig for %v.\n", remainigTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v ticket for %v %v.", userTickets, firstName, lastName)
	fmt.Println("—————————————————")
	fmt.Printf("Sending ticket:\n%v \nto email adress %v.\n", ticket, email)
	fmt.Println("—————————————————")

} 