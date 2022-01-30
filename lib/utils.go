package utils

import (
	"fmt"
	"log"
)

// global variables
var accounts = make(map[string]string)

func InitAccounts() map[string]string {
	accounts["Test"] = "test"
	return accounts
}

func GetLoginData(id string) string {
	if accounts[id] == "" {
		log.Printf("No user with Id: %s", id)
		userNotFound := fmt.Sprintf("No user with Id: %s", id)
		return userNotFound
	}
	userFound := fmt.Sprintf("Welcome, %s", id)
	return userFound
}

func FetchTemplate(tpl string) string {
	return tpl
}
