package main

import "fmt"

func main() {
	confrenceName := "Go Confrence" // automatically defines type and treat is as a variable
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Println("Welcome to", confrenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still availiable")
	fmt.Println("You can book your tickets here")

	//var booking = [conferenceTickets]string{}
	booking := []string{}
	for {
		var firstName string
		var LastName string
		var email string
		var userTickets uint
		fmt.Println("Enter your First Name :")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Last Name :")
		fmt.Scan(&LastName)
		fmt.Println("Enter your Email Address :")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you want")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Println("Its more than remaining tickets we only have", remainingTickets, "tickets left so you cannot book", userTickets, " tickets.")
			fmt.Println("Plaese try again !")
			continue
		}
		remainingTickets = remainingTickets - userTickets
		//booking[0] = firstName + LastName
		booking = append(booking, email)
		fmt.Println("You will be getting your tickets on your  email")
		fmt.Println(remainingTickets, " tickets remaining for", confrenceName)
		fmt.Println("Booking till now", booking)
		if remainingTickets == 0 {
			//end program
			fmt.Println("Conferenece is booked out !")
			break
		}
	}

}
