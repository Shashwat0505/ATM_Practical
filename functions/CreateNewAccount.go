package functions

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CreateNewAccount() {
	fmt.Println("Please enter username")
	var username string
	fmt.Scan(&username)
	min := 1000000000
	max := 9999999999
	rand.Seed(time.Now().UnixNano())
	accNoInt := (rand.Intn(max-min) + min)
	accNo := strconv.Itoa(accNoInt)
	fmt.Println("Your account no is:-", accNo)

	min1 := 1000
	max1 := 9999
	pinNo := (rand.Intn(max1-min1) + min1)
	accPinNo := strconv.Itoa(pinNo)
	fmt.Println("PIN for your new account is:-", accPinNo)

label4:
	fmt.Println("Please enter amount of balance you want to add to account in ₹")
	var inputAmount string
	fmt.Scan(&inputAmount)
	inputAmount = strings.ReplaceAll(inputAmount, ",", "")
	if !IsValid(inputAmount) {
		fmt.Println("Invalid amount entered")
		goto label4
	}

	Users = append(Users, User{UserName: username, PIN: accPinNo, AccountNo: accNo, Balance: inputAmount})

	return

}
