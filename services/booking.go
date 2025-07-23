package services

import (
	"fmt"
	"miniBookingApp/model"
)

var bookings []model.User

func AddBookings(user model.User) {
	bookings = append(bookings, user)
	fmt.Printf("\n%v tickets booked for %v %v", user.Tickets, user.FirstName, user.LastName)
}

func ShowBookings() {
	fmt.Println("\nAll bookings:")
	if len(bookings) == 0 {
		fmt.Println("No bookings yet.")
		return
	}

	for i, user := range bookings {
		fmt.Printf("%v. %v %v (%v tickets)\n", i+1, user.FirstName, user.LastName, user.Tickets)
	}
}
