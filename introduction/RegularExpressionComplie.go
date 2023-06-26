package main
import (
		"fmt"
		"regexp"
	)
func main(){
	str := "cat mat rat pat fat"
	reg, _  := regexp.Compile("([cmrpf]at)")
	fmt.Println("Match string", reg.MatchString(str))
	fmt.Println("Find String", reg.FindString(str))
	fmt.Println("Find the string Index", reg.FindStringIndex(str))
	fmt.Println("Find all match string", reg.FindAllString(str, -1))
	fmt.Println("Find First 2nd string", reg.FindAllString(str, 2))
	fmt.Println("All Submatch index", reg.FindAllStringSubmatchIndex(str, -1))
	fmt.Println("Replace the string", reg.ReplaceAllString(str, "Dog"))
}