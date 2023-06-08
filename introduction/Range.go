package main
import (
	"fmt"
)
var pl = fmt.Println
func main(){
	nArry := []int{1,2,4,6}
	for _, num :=range nArry{
		pl(num);
	}
}