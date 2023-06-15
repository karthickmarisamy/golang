package main
import (
	"fmt"
)
type customer struct{
	name string
	age int
	salary float64
}

func getCustomer(sCustom customer){
	fmt.Printf("Customer name is %s\n Customer Age is %d\n Customer salary %.2f", sCustom.name, sCustom.age, sCustom.salary)
}

func addNewCustomer(sCustom *customer, salary float64){
	sCustom.salary = salary
}

func main(){
	var sCustomer customer
	sCustomer.name = "John"
	sCustomer.age = 34
	sCustomer.salary = 4.3
	getCustomer(sCustomer)
	addNewCustomer(&sCustomer, 123.3312)
	oCustomer := customer{"Johny",16,123.3123}
	fmt.Println(oCustomer);
}
