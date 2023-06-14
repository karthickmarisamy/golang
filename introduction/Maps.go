package main
import (
	"fmt"
)

func main(){
	var iMap map[string]string
	iMap = make(map[string]string)
	iMap["name"] = "Johny"
	iMap["email"] = "email"
	iMap["job"] = "programmer"
	
	xMap := map[int]string{1:"Test", 2:"Test1", 3:"Test3"}

	_, ok := xMap[4]
	fmt.Println("The invalid index", ok)

	for k, v := range iMap{
		fmt.Printf("The key is %s and value is %s\n", k, v)
	}

	delete(xMap, 3)

	fmt.Println(xMap)
}
