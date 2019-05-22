package bank

import (
	"bufio"
	"fmt"
	"os"
)

//User is a struct that defines what a fields a User has
type User struct {
	Name     []string
	Email    string
	Password string
	Pin      int
	Balance  int
}

func (u *User) register(email string, password string, pin int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(scanner)
}

func (u *User) transact() {

}

func (u *User) login(email string, password string, pin int) {

}
