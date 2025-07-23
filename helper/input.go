package helper

import(
	"fmt"
	"strings"
	"miniBookingApp/model"
)

func GetUserInput() model.User {
	var firstName,lastName,email string
	var tickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want: ")
	fmt.Scan(&tickets)

	return model.User{
		FirstName : firstName,
		LastName : lastName,
		Email : email,
		Tickets : tickets,

	}
}

func ValidUserInput ( user model.User ) bool {
	isValidName := len(user.FirstName) >= 2 && len(user.LastName)>=2
	isValidEmail := strings.Contains(user.Email,"@")
	isValidTickets := user.Tickets>0

	if !isValidName {
		fmt.Println("Invalid Input!\nFirstname and Lastname should be atleast 2 characters long")
	}

	if !isValidEmail {
		fmt.Println("Email must contain '@'.")
	}

	if !isValidTickets {
		fmt.Println("You must book atleast one ticket.")
	}

	return isValidName && isValidEmail && isValidTickets
}