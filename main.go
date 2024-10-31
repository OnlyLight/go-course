package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTicket uint = 5

func main() {
	var remainingTickets uint = 5
	conferenceName := "Go Conference"
	bookings := make([]helper.UserData, 0)

	for remainingTickets > 0 {
		helper.GreetUsers(conferenceName, conferenceTicket, remainingTickets)

		userName, firstName, userTickets := helper.InputValue()

		if userTickets > remainingTickets {
			fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("\nRemaining ticket: %d", remainingTickets)

		userData := helper.UserData{
			UserName:    userName,
			FirstName:   firstName,
			UserTickets: userTickets,
		}

		bookings = append(bookings, userData)
		fmt.Printf("\nThank u %v\n", firstName)

		helper.Wg.Add(1)
		go helper.SleepToSend()
		helper.Wg.Wait()

		if remainingTickets == 0 {
			fmt.Println("Sold out")
			break
		}

	}

	fmt.Printf("\nThank you for bookings %v tickets.", conferenceTicket-remainingTickets)
	fmt.Printf("\nList First name: %v\n", helper.GetFristName(bookings))
	fmt.Printf("\nThese are all bookings %v\n", bookings)
}
