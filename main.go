package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"golang.org/x/term"
)

const accbalancefile = "balance.txt"

func Readbalance() float64 {
	data, _ := os.ReadFile(accbalancefile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func Writebalance(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accbalancefile, []byte(balanceText), 0644)
}

var AccountBalance = Readbalance()
var UserId, Password string

func main() {
	fmt.Println("Welcome to Go Bank")
	fmt.Print("Please enter your user id:")
	fmt.Scanln(&UserId)
	fmt.Print("Please enter your password:")
	//fmt.Scan(&Password)

	//fmt.Print("Please enter your password: ")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	Password = string(bytePassword)
	//fmt.Println()

	if UserId == "admin" && Password == "admin" {
		fmt.Println("Welcome to Go_Bank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	} else {
		fmt.Println("Invalid User Id or Password")
		return
	}

	var choice int
	fmt.Print("Please enter your choice:")
	fmt.Scan(&choice)
	fmt.Println("Your choice is:", choice)

	switch choice {
	case 1:
		fmt.Println("Your account balance is:", AccountBalance)
	case 2:
		var depositAmount float64
		fmt.Print("Please enter the amount to be deposit: ")
		fmt.Scan(&depositAmount)
		if depositAmount <= 0 {
			fmt.Println("Invalid amount")
		} else {
			AccountBalance += depositAmount
			fmt.Println("Your new account balance is:", AccountBalance)
		}
		Writebalance(AccountBalance)
	case 3:
		fmt.Print("please enter the amount to be withdrawn: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > AccountBalance {
			fmt.Println("Insufficient funds")
		} else {
			AccountBalance -= withdrawAmount
			fmt.Println("Your new balance is", AccountBalance)
		}
		Writebalance(AccountBalance)
	default:
		fmt.Println("Thank you for using Go Bank")
	}
}
