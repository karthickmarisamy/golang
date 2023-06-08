package main
import ("fmt")
var pl = fmt.Println

func getTwo(x int) (int, int){
	return x + 1, x + 2
}

func sumNumber(x int, y int) int{
	return x + y;
}

func getQuotient(x float64, y float64) (ans float64, err error){
	if y == 0{
		return 0, fmt.Errorf("You can't divide by 0")
	} else {
		return x / y , nil
	}

}
func main(){
	pl(getTwo(2))
	pl(sumNumber(1, 2))
	pl(getQuotient(1, 1))
}