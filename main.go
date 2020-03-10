package main

import (
	"fmt"

	"github.com/ehddnr301/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("superman")
	fmt.Println(account)
}
