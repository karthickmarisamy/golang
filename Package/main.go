package main
import (
	"fmt"
	stuff "sample/project/mypackage"
	"reflect"
)

func main(){
	fmt.Println(stuff.Name);
	intArry := []int{1,2,3,4,5}
	strArry := stuff.IntArrytoStrArry(intArry)
	fmt.Println(strArry)
	fmt.Println(reflect.TypeOf(strArry))


}