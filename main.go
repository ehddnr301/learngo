package main

import (
	"fmt"

	"github.com/ehddnr301/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	word := "Hello"
	definition := "Hi to someone"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println("Found ", word, "/ Definition:", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
