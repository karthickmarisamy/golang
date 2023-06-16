package main
import (
	"fmt"
	stuff "sample/project/mypackage"
	"reflect"
	"log"
)

func main(){
	fmt.Println(stuff.Name);
	intArry := []int{1,2,3,4,5}
	strArry := stuff.IntArrytoStrArry(intArry)
	fmt.Println(strArry)
	fmt.Println(reflect.TypeOf(strArry))

	date := stuff.Date{}

	var err = date.SetDay(1)
	if err != nil{
		log.Fatal(err)
	}

	err = date.SetMonth(1)
	if err != nil{
		log.Fatal(err)
	}

	err = date.SetYear(1991)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("The Current data is %d-%d-%d\n", date.Date(), date.Month(), date.Year())
}