package bank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//User is a struct that defines what a fields a User has
type User struct {
	Name     string
	Email    string
	Password string
	Pin      int
	Balance  int
}

func (u *User) start() {
	fmt.Println(" Please login or register by typing login or register")
	strTok := ""

	switch strings.ToUpper(strTok) {
	case "register":
		u.register(u.Email, u.Password, u.Pin)
	case "login":
		u.login(u.Email, u.Password, u.Pin)
	}
}

func (u *User) register(email string, password string, pin int) *User {
	fmt.Println("please enter an email and a password")
	u = &User{
		Email:    email,
		Password: password,
	}
	u.Save(email, password)
	return u
}

func (u *User) userCommands(command string) (int int) {
	balance := u.Balance

	return balance
}

func (u *User) login(email string, password string, pin int) {

	if u.Password != password {
		fmt.Println("incorrect password")
	}
	if u.Email != email {
		fmt.Println("incorrect username")
	} else {
		if u.Email == email && u.Password == password {
			fmt.Println("Login successful")
			//	u.Load()
		}
	}
}

//Save saves a user's account information
func (u *User) Save(user string, password string) {
	u.Email = user
	data, err := json.Marshal(u)
	if err != nil {
		log.Fatalln("Failed to save")
	}
	err = ioutil.WriteFile(user, data, 0644)
	if err != nil {
		log.Fatalln("Failed to write.")
	}
}

//Load loads a user's account information
func (u *User) Load(user string) {
	dat, err := ioutil.ReadFile(user)

	err = json.Unmarshal(dat, u)
	if err != nil {
		log.Println("no data file")
	}
}
