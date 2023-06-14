package main
import(
	"fmt"
)

type MyConstraint interface{
	int | float64
}

func getSumGen [T MyConstraint] (x T, y T) T{
	return x + y
}

func main(){
	fmt.Println("The sum of two values", getSumGen(4,5))
	fmt.Println("The sum of two values", getSumGen(4.0,5.0))
}