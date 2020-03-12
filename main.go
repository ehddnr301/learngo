package main

import (
	"fmt"
	"log"

	"github.com/ehddnr301/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("superman")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
