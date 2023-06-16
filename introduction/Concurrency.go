package main
import ("fmt"
"time"
)
func print1To15(){
	for i := 0; i < 15; i++{
		fmt.Println("Func 1", i)
	}
}

func print1To10(){
	for i := 0; i < 10; i++{
		fmt.Println("Func 2", i)
	}
}

func main(){
	go print1To10()
	go print1To15()
	time.Sleep(2 * time.Second)
}