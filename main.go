package main

import (
	"web/database"
	"web/register"
	"web/login"
	"web/download"
	"os"
)



func main() {
	status := database.PingDb()
	action := os.Args[1]
	if action == "register" {
		if status == true {
			register.RegisterUser()
		}
	} else if action == "login" {
		if status == true {
			login.LogIn()
		}
	} else if action == "download" {
		if status == true {
			download.DownloadConfFile()
		}
	}
}
