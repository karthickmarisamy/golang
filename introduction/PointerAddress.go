package main
import ("fmt")
var pl = fmt.Println

func changeArryVal(myPtr *[4]int){
	for i:=0; i < 4 ; i++{
		myPtr[i]*=2
	}
}
func main(){
	f := 10;
	var ptrVar *int = &f;
	pl("The pointer address", ptrVar)
	pl("The pointer value", *ptrVar)

	var fArry = [4]int{1,2,3,4}
	pl("Befor changing the values in array", fArry)
	changeArryVal(&fArry)
	pl("After changing the values in array", fArry)

}