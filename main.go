package main

import (
	"fmt"

	"github.com/ehddnr301/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Delete(baseWord)
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
