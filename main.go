package main

import (
	"fmt"

	"github.com/ehddnr301/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
