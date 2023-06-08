package main
import (
	"fmt"
)
var pl = fmt.Println
func main()  {
	sArry := make([]string, 5)
	sArry[0] = "first"
	sArry[1] = "second"
	sArry[2] = "third"
	sArry[3] = "fourth"
	sArry[4] = "fifth"
	for _, val := range sArry{
		pl("The array values are ", val);
	}

	pl("Array Slice from the Beginning")
	sl1 := sArry[:3]
	pl(sl1)
	pl("Array Slice from the Last")
	sl2 := sArry[2:]
	pl(sl2)

	sl2 = append(sl2, "new")
	pl(sl2)
}