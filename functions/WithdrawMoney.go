package functions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func WithdrawMoney(u *User) {
	fmt.Println("Your current balance is ₹", u.Balance)
label6:

	fmt.Println("!---Decimal values shold not be entered---!")
	fmt.Println("Enter the amount in multiple of 500 and greater than 500 to withdraw:-")
	var withdrawAmont string
	fmt.Scan(&withdrawAmont)

	if !IsValid(withdrawAmont) {
		fmt.Println("Invalid amount!!Please try again")

	label7:
		fmt.Println("Do you want to continue??")
		fmt.Println("Press y for continue:-")
		fmt.Println("Press n for exit:-")
		var input1 string
		fmt.Scan(&input1)
		input1 = strings.ToLower(input1)
		switch input1 {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***")
			time.Sleep(time.Millisecond)
			os.Exit(0)

		default:
			fmt.Println("Invalid input")
			goto label7
		}

	}
	if CheckforDot(withdrawAmont) {
		fmt.Println("Invalid amount!!Please try again")
		goto label6

	}
	tempBalance, _ := strconv.ParseFloat(u.Balance, 64)
	withdrawcash, _ := strconv.ParseFloat(withdrawAmont, 64)

	if withdrawcash > 20000 {
		fmt.Println("!!!You cant withdraw more than 20,000 in a single day!!!")
	label12:
		fmt.Println("Do you want to continue??")

		fmt.Println("Press y for continue:-")
		fmt.Println("Press n for exit:-")
		var input3 string
		fmt.Scan(&input3)
		input3 = strings.ToLower(input3)
		switch input3 {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***")
			time.Sleep(time.Millisecond)
			os.Exit(0)

		default:
			fmt.Println("Invalid input")
			goto label12
		}

	}

	if int(withdrawcash)%500 != 0 {
		fmt.Println("!!!Amount shold be in multiple of 500!!!")
	label13:
		fmt.Println("Do you want to continue??")

		fmt.Println("Press y for continue:-")
		fmt.Println("Press n for exit:-")
		var input3 string
		fmt.Scan(&input3)
		input3 = strings.ToLower(input3)
		switch input3 {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***")
			time.Sleep(time.Millisecond)
			os.Exit(0)

		default:
			fmt.Println("Invalid input")
			goto label13
		}

	}

	if !IsValidWithdrawAmount(withdrawcash, tempBalance) {

		fmt.Println("!---Insufficient Balance---!")

	label8:
		fmt.Println("Do you want to continue??")

		fmt.Println("Press y for continue:-")
		fmt.Println("Press n for exit:-")
		var input2 string
		fmt.Scan(&input2)
		input2 = strings.ToLower(input2)
		switch input2 {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***")
			time.Sleep(time.Millisecond)
			os.Exit(0)

		default:
			fmt.Println("Invalid input")
			goto label8
		}
	} else {

		tempBalance = tempBalance - withdrawcash
		u.Balance = fmt.Sprintf("%v", tempBalance)
		fmt.Println("Now your current balance is:-₹", u.Balance)

	label9:
		fmt.Println("Do you want to continue?")
		fmt.Println("Press y for yes")
		fmt.Println("Press n for no")

		var choose string
		fmt.Scan(&choose)
		choose = strings.ToLower(choose)

		switch choose {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***!")
			time.Sleep(time.Millisecond)

			os.Exit(0)

		default:
			fmt.Println("Incorrect input!Please try again")
			goto label9

		}
	}
}
