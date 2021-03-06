package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/SaucySeadweller/bankingProgram/bank"
)

func main() {
	fmt.Println("Please login or register by typing login or register")
	cmd := GetString()
	executeCmd(cmd)
}
func executeCmd(cmd string) {

	switch strings.ToLower(cmd) {
	default:
		log.Println("command not found", cmd)
	case "register":
		u := bank.Register(GetNamePass())
		log.Println(u)
		fmt.Println("Please make an intial deposit.")
		u.Deposit(GetFloat())
		fmt.Println(u.Balance)
		u.Save()

	case "login":
		u, err := bank.Login(GetNamePass())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Welcome:", u.Name)

		for {
			fmt.Println("What do would you like to do?")
			cmd := GetString()
			switch strings.ToLower(cmd) {
			default:
				log.Println("command not found", cmd)
			case "log":
				fmt.Println(u.ViewTransactions())
			case "balance":
				fmt.Println(u.Balance)
			case "deposit":
				fmt.Println("enter an amount to deposit")
				u.Deposit(GetFloat())
				fmt.Println(u.Balance)
			case "withdrawl":
				fmt.Println("enter an amount to withdrawl")
				u.Withdrawl(GetFloat())
				fmt.Println(u.Balance)
			}
			u.Save()
		}
	}
}

//GetNamePass checks login credentials
func GetNamePass() (string, string) {
	var name string
	var password string
	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name)
	fmt.Println("What is your password?")
	fmt.Scanf("%s", &password)
	return name, password
}

//GetString gets commands from the user
func GetString() string {
	cmd := ""
	_, err := fmt.Scanf("%s", &cmd)
	if err != nil {
		panic(err)
	}
	return cmd
}

//GetFloat gets the transactiion amount
func GetFloat() float64 {
	amount := 0.00
	_, err := fmt.Scanf("%f", &amount)
	if err != nil {
		panic(err)
	}
	return amount
}
