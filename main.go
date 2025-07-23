package main

import (
	"fmt"
	"miniBookingApp/helper"
	"miniBookingApp/model"
	"miniBookingApp/services"
	"miniBookingApp/utils"
	"sync"
)

const Tickets = 100

var remainingTickets = 100

func printInfo() {
	fmt.Printf("Total tickets: %v\n", Tickets)
	fmt.Printf("Remaining Tickets available: %v\n", remainingTickets)
}

func main() {
	var wg sync.WaitGroup
	printInfo()

	for remainingTickets > 0 {
		user := helper.GetUserInput()

		if helper.ValidUserInput(user) {
			if user.Tickets <= uint(remainingTickets) {
				services.AddBookings(user)
				remainingTickets -= int(user.Tickets)
				fmt.Printf("\nRemaining tickets: %v",remainingTickets)

				wg.Add(1)
				go func(u model.User) {
					defer wg.Done()
					utils.SendConfirmation(u)
				}(user)
			} else {
				fmt.Printf("Sorry, only %v tickets remaining!\n", remainingTickets)
			}
		} else {
			fmt.Println("Please try again with valid input.")
		}

		services.ShowBookings()
	}

	wg.Wait()
	fmt.Println("All tickets booked! The event is now fully booked.")
}
