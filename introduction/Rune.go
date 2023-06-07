package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main(){
	strVal := "Hello"
	pl("Total number of character in a string", utf8.RuneCountInString(strVal))
	for i, runes := range strVal{
		fmt.Printf("%d, %#U, %c\n", i, runes, runes)
	}
}