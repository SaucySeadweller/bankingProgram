package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SaucySeadweller/bankingProgram/bank"
)

func main() {
	var timeStamp time.Time
	var amount float64
	fmt.Println("Please login or register by typing login or register")
	cmd := GetString()
	switch strings.ToLower(cmd) {
	default:
		log.Println("command not found", cmd)
	case "register":
		u := bank.Register(GetEmailPass())
		log.Println(u)
		u.Save()
	case "login":
		u, err := bank.Login(GetEmailPass())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Welcome: ", u.Name)
		fmt.Println("What do you want to do?")
		cmd := GetString()
		switch strings.ToLower(cmd) {
		default:
			log.Println("command not found", cmd)
		case "log":
			u.TransactionLog()
		case "balance":
			fmt.Println(u.Balance)
		case "deposit":

			u.Deposit(amount)
			fmt.Println(u.Balance)
		case "withdrawl":
			u.Withdrawl(amount)
			fmt.Println(u.Balance)
		case "history":
			u.NewTransaction(amount, timeStamp)
		}
	}
}

//GetEmailPass checks login credentials
func GetEmailPass() (string, string) {
	var email string
	var password string
	fmt.Println("What is your email?")
	fmt.Scanf("%s", &email)
	fmt.Println("What is your password?")
	fmt.Scanf("%s", &password)
	return email, password

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
