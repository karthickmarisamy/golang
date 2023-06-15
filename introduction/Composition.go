package main
import("fmt")

type contact struct{
	fName string
	lName string
	age int
}
type business struct{
	companyName string
	address string
	contact
}

func (b business)Info(){
	fmt.Printf("The first name is %s\n The last name is %s\n The age is %d\n The company name is %s\n The address is %s\n The contact is %s", b.contact.fName, b.contact.lName, b.contact.age, b.companyName, b.address, b.contact)
}

func main(){
	cContact := contact{"James", "Wang", 23}
	cInfo := business{
		"Startup",
		"Global",
		cContact,
	}
	cInfo.Info()
}