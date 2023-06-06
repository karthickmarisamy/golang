package main
import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main(){
	pl(reflect.TypeOf(23))
	pl(reflect.TypeOf(2.3))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf(true))

	// casting the variable type

	// string to integer
	var stringVar = "1234"
	convertVal, err := strconv.Atoi(stringVar)
	pl(convertVal, err, reflect.TypeOf(convertVal))

	//Integer to Ascii

	var intVar = 1234
	pl(strconv.Itoa(intVar))

	//string to float

	var stringFloat = "2.3"
	if convFloat, err := strconv.ParseFloat(stringFloat, 64); err==nil{
		pl(convFloat);
	}
}
