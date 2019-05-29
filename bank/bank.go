package bank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

//User is a struct that defines what a fields a User has
type User struct {
	Name         string
	Email        string
	Password     string
	Balance      float64
	Transactions []Transaction
}

//Transaction is a struct that contains the amount and time of a transaction
type Transaction struct {
	amount float64
	date   time.Time
}

//Register registers a new user in the system
func Register(email string, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func (u *User) accountBalance() float64 {
	return u.Balance
}

//Deposit deposits an amount into the user's account
func (u *User) Deposit(amount float64) float64 {
	u.NewTransaction(amount, time.Now())
	u.Balance -= amount
	return u.Balance
}

//Withdrawl withdraws an amount from the user's account
func (u *User) Withdrawl(amount float64) float64 {
	u.NewTransaction(amount, time.Now())
	u.Balance += amount
	return u.Balance
}

//Login allows the user to login to their account
func Login(email string, password string) (*User, error) {
	u := &User{}
	if !u.Load(email) {
		return nil, fmt.Errorf("invalid email/password")
	}
	if u.Password != password || u.Email != email {
		return nil, fmt.Errorf("incorrect email/password")
	}
	return u, nil
}

//TransactionLog brings up a log of the users past transactions
func (u *User) TransactionLog() {
	fmt.Println(u.Transactions)
}

//NewTransaction keeps track of the transactions a user makes
func (u *User) NewTransaction(amount float64, when time.Time) {
	u.Transactions = append(u.Transactions, Transaction{amount, when})
	fmt.Println("transaction made:", u.Transactions[len(u.Transactions)-1])
}

//Save saves a user's account information
func (u *User) Save() {
	data, err := json.Marshal(u)
	if err != nil {
		log.Fatalln("Failed to save")
	}
	err = ioutil.WriteFile(u.Email, data, 0644)
	if err != nil {
		log.Fatalln("Failed to write.")
	}
}

//Load loads a user's account information
func (u *User) Load(user string) bool {
	dat, err := ioutil.ReadFile(user)
	if err != nil {
		return false
	}
	err = json.Unmarshal(dat, u)
	if err != nil {
		log.Fatalln("no data file")
	}
	return true
}
