package bank

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//User is a struct that defines what a fields a User has
type User struct {
	Name     []string
	Email    string
	Password string
	Pin      int
	Balance  int
	Commands []Command
}

//Command is struct containing actions a user can take
type Command struct {
	checkBalance string
	deposit      string
	withdraw     string
}

func main() {
	fmt.Println(" Please login or register")

}

func (u *User) register(email string, password string, pin int) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(scanner)

}

func (u *User) transact(command string) (int int) {
	balance := u.Balance

	return balance
}

func (u *User) login(email string, password string, pin int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(scanner)
}

//Save saves a user's account information
func (u *User) Save(user string) {
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
