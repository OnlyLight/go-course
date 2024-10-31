package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var Wg = sync.WaitGroup{}

type UserData struct {
	UserName    string
	FirstName   string
	UserTickets uint
}

func InputValue() (string, string, uint) {
	var userName string
	var userTickets uint
	var firstName string

	fmt.Printf("\nInput username: ")
	fmt.Scan(&userName)
	fmt.Printf("Username: %s", userName)

	fmt.Printf("\nInput firstName: ")
	fmt.Scan(&firstName)
	fmt.Printf("firstName: %s", firstName)

	fmt.Printf("\nInput tickets: ")
	fmt.Scan(&userTickets)
	fmt.Printf("userTickets: %d", userTickets)

	return userName, firstName, userTickets
}

func GreetUsers(conferenceName string, conferenceTicket uint, remainingTickets uint) {
	fmt.Println("Welcome to", conferenceName)
	fmt.Println("We have", conferenceTicket, "tickets and", remainingTickets, "still available")
}

func GetFristName(bookings []UserData) []string {
	temps := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking.FirstName)
		temps = append(temps, names[0])
	}

	return temps
}

func SleepToSend() {
	time.Sleep(10 * time.Second)
	fmt.Println("***************")
	fmt.Println("Sending........")
	fmt.Println("***************")
	Wg.Done()
}
