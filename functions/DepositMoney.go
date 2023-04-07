package functions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DepositMoney(u *User) {
	fmt.Println(u.UserName)
	fmt.Println("Your current balance is ₹", u.Balance)
label3:
	fmt.Println("Enter the amount you want to deposit:-₹")
	var input string
	fmt.Scan(&input)

	if !IsValid(input) {
		fmt.Println("Invalid amount!!Please try again")
		goto label3

	}

	tempBalance, _ := strconv.ParseFloat(u.Balance, 64)

	depositAmount, _ := strconv.ParseFloat(input, 64)
	if int(depositAmount)%500 != 0 {
		fmt.Println("!!!Amount shold be in multiple of 500 and greater than 100!!!")

		fmt.Println("Do you want to continue?")
		fmt.Println("Press y for yes")
		fmt.Println("Press n for no")

		var choose string
		fmt.Scan(&choose)
		choose = strings.ToLower(choose)
	label14:
		switch choose {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***!")
			time.Sleep(time.Millisecond)

			os.Exit(0)

		default:
			fmt.Println("Incorrect input!Please try again")
			goto label14

		}

	}

	if depositAmount > 50000 {
		fmt.Println("You can't enter this large amount ₹", depositAmount)
		fmt.Println("!--Please reduce this amount---!")

		fmt.Println("Do you want to continue?")
		fmt.Println("Press y for yes")
		fmt.Println("Press n for no")

		var choose string
		fmt.Scan(&choose)
		choose = strings.ToLower(choose)
	label10:
		switch choose {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***!")
			time.Sleep(time.Millisecond)

			os.Exit(0)

		default:
			fmt.Println("Incorrect input!Please try again")
			goto label10

		}

	}
	if depositAmount < 100 {
		fmt.Println("You can't enter this  amount ₹", depositAmount)
		fmt.Println("!--Minimum amount should be greater than 100---!")

		fmt.Println("Do you want to continue?")
		fmt.Println("Press y for yes")
		fmt.Println("Press n for no")

		var choose string
		fmt.Scan(&choose)
		choose = strings.ToLower(choose)
	label11:
		switch choose {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***!")
			time.Sleep(time.Millisecond)

			os.Exit(0)

		default:
			fmt.Println("Incorrect input!Please try again")
			goto label11

		}

	}
	tempBalance += depositAmount

	u.Balance = fmt.Sprintf("%v", tempBalance)

	fmt.Println("Now your balance is:-₹", u.Balance)
label2:
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
		goto label2

	}

}
