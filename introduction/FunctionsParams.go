package main
import ("fmt")

func arrArgument(number ...int) int{
	sum := 0
	for _, num := range number{
		sum += num
	}
	return sum;
}

func arraySum(arry []int) int{
	sum := 0
	for _, num := range arry{
		sum += num
	}
	return sum;
}


func main(){
	arrys := []int{1,24,5,6,6}
	fmt.Println(arrArgument(1,2,4,5,6))
	fmt.Println(arraySum(arrys))
}