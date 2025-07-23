package utils

import (
	"fmt"
	"miniBookingApp/model"
	"time"
)

func SendConfirmation(user model.User){
	time.Sleep(time.Second)
	fmt.Printf("\nSending confirmation email to %v for %v tickets.\n",user.Email,user.Tickets)
}