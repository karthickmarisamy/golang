// need to run like go build Args.go
// go ./Args 1 2 4
package main
import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	fmt.Println(os.Args)
	iArgs := []int{}
	args := os.Args[1:]
	for _, i := range args{
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, i := range iArgs{
		if i > max {
			max = i
		}
	}

	fmt.Println("The Max Value is ", max)
}