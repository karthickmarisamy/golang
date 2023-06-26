package main
import (
		"fmt"
		"regexp"
	)

func main(){
	str := "This is the string from go lang"
	match, _ := regexp.MatchString("(go[^ ]?", str)
	fmt.Println(match)
}