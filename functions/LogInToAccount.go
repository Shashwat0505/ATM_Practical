package functions

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	UserName  string
	PIN       string
	AccountNo string
	Balance   string
}

var Users = []User{
	{UserName: "Shashwat", PIN: "5522", AccountNo: "555", Balance: "5500"},
}

func LogInToAccount() (User, error) {

	fmt.Println("Please enter accountno:-")
	var acno string
	fmt.Scan(&acno)

	fmt.Println("Please enter pin:-")
	var pin string
	fmt.Scan(&pin)

	for _, val := range Users {

		if val.AccountNo == acno && val.PIN == pin {

			fmt.Println("!***Welcome ", val.UserName, "***!")
			fmt.Println(strings.Repeat("*", 26))

			return val, nil
		}
	}

	fmt.Println("Sorry invalid account-number or pin number")
	temp := User{}
	return temp, errors.New("Invalid credentials")

}
