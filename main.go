package main

import (
	"fmt"

	"github.com/ehddnr301/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("superman")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
