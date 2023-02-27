package login

import (
	"fmt"
	"web/api"
	"syscall"
	"golang.org/x/term"
	"web/hasher"
	"gitlab.com/david_mbuvi/go_asterisks"
	"os"
)

func LogIn() (bool, string){
	var user string
	fmt.Println("User login")
	fmt.Print("Enter username: ")
	fmt.Scan(&user)
	fmt.Print("Enter password: ")
	passwordOne, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	fmt.Print("Enter security token: ")
	input, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
	token := string(input)
	password := string(passwordOne)
	hash := checker.HashPassword(password)
	fmt.Println("")
	exists := api.Apiuser(user, hash, token)
	if exists == true {
		fmt.Println("Valid credentials")
	} else {
		fmt.Println("Invalid credentials")
	}
	return exists,token
}
