package register

import (
	"fmt"
	"os"
	"syscall"
	"web/hasher"
	"web/database"

	"golang.org/x/term"
)

func RegisterUser(){
	var user string
	fmt.Println("User registration")
	fmt.Print("Enter username: ")
	fmt.Scan(&user)
	fmt.Print("Enter password: ")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	fmt.Print("Repeat the password: ")
	bytePasswordTwo, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	passwordOne := string(bytePassword)
	passwordTwo := string(bytePasswordTwo)
	if passwordOne == passwordTwo {
		hash := checker.HashPassword(passwordOne)
		var base64File string
		fmt.Print("Enter kubeconfig in base64: ")
		fmt.Scan(&base64File)
		database.CreateUser(user, hash, base64File)
	} else {
		fmt.Println("Password does not match")
		os.Exit(0)
	}

}