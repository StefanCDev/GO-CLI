package main

import (
	"fmt"
    "strings"
)





func main() {
	conferenceName := "Go Conference"
	const  conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string



	fmt.Printf("Welcome to %v booking application\n", conferenceName )
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
	
		fmt.Println("Enter number if tickets")
		fmt.Scan(&userTickets)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isVAlidEmail bool = strings.Contains(email, "@")
		var isValidTicket bool = userTickets > 0 && userTickets <= remainingTickets
		// validate user input

		if isValidName && isVAlidEmail && isValidTicket {
		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)
	
		
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)


		firstNames := []string{}
		for _, booking := range  bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
			 
		}
		fmt.Printf("The first names of bookings are :  %v\n", firstNames)



	
		if remainingTickets == 0 {
		// end program
		fmt.Println("Our conference is booked out. Come back next year.")
		     break
		}
		} else {
			if !isValidName{
				fmt.Printf("First name and last name you entered is too short.\n")
			}
			if !isVAlidEmail{
				fmt.Printf("Email address you entered doesn`t contain '@'.\n")
			}
			if !isValidTicket{
				fmt.Printf("Number of tickets you entered is invalid.\n")
			}
			fmt.Printf("Booking is not valid. Please try again.\n") 

		}
	}
	
}
