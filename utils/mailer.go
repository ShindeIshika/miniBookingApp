package utils

import (
	"fmt"
	"miniBookingApp/model"
	"time"
)

func SendConfirmation(user model.User){
	time.Sleep(10*time.Second)
	fmt.Printf("\nConfirmation email sent to %v for %v tickets.\n",user.Email,user.Tickets)
}