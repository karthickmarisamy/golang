package main
import ("fmt")
var pl = fmt.Println

func changeVal(myPtr *int){
	*myPtr = 12
}

func main(){
	f := 4
	pl("before change the value", f)
	changeVal(&f)
	pl("after change the value", f)
}