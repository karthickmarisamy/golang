package main
import ("fmt")
func num1(channel chan int){
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
}

func num2(channel chan int){
	channel <- 5
	channel <- 6
	channel <- 7
}

func main(){
	channel1 := make(chan int)
	channel2 := make(chan int)
	go num1(channel1)
	go num2(channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
}