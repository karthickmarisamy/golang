package main
import ("fmt")

func funFun (f func(int, int) int, x, y int){
	fmt.Println(f(x, y))
}

func sampleSum(x , y int) int{
	return x + y
}

func main(){
	intSum := func(x, y int)int{
		return x + y
	}
	fmt.Println("The sum value is ", intSum(2, 5))
	testingVal := 1
	fmt.Println("The testing val is", testingVal)
	changeVal := func(){ testingVal += 1 }
	changeVal()
	fmt.Println("The testing val is", testingVal)
//	fmt.Println(funFun(sampleSum, 5,5))
	funFun(sampleSum, 5,5)

}