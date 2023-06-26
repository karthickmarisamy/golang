package main
import ("fmt")

func factorial(num int)int{
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}

func main(){
	fmt.Println("The factorial value of 4 is ", factorial(4));
}