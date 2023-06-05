package main
import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

func main()  {
	pl("Please enter ur liks")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil{
		pl("Your lik is ", name)
	}else{
		pl(err)
	}
}