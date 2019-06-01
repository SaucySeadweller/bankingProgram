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
	Password     string
	Balance      float64
	Transactions []Transaction
}

//Transaction is a struct that contains the amount and time of a transaction
type Transaction struct {
	Amount float64
	Date   time.Time
}

//Register registers a new user in the system
func Register(Name string, password string) *User {

	return &User{
		Name:     Name,
		Password: password,
	}
}

func (u *User) accountBalance() float64 {
	return u.Balance
}

//Deposit deposits an amount into the user's account
func (u *User) Deposit(amount float64) float64 {
	u.NewTransaction(amount, time.Now())
	u.Balance += amount

	return u.Balance
}

//Withdrawl withdraws an amount from the user's account
func (u *User) Withdrawl(amount float64) (int, float64) {
	u.NewTransaction(amount, time.Now())
	u.Balance -= amount
	return 0, u.Balance
}

//ViewTransactions prints out prior tansactions and their time stamps
func (u *User) ViewTransactions() string {
	str := ""
	for _, t := range u.Transactions {
		str += fmt.Sprintf("%f %s \n", t.Amount, t.Date.Format(time.ANSIC))
	}
	return str
}

//Login allows the user to login to their account
func Login(Name string, password string) (*User, error) {
	u := &User{}
	if !u.Load(Name) {
		return nil, fmt.Errorf("invalid Name/password")
	}
	if u.Password != password || u.Name != Name {
		return nil, fmt.Errorf("incorrect Name/password")
	}
	return u, nil
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
	err = ioutil.WriteFile(u.Name, data, 0644)
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
