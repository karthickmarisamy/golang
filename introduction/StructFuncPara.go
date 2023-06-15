package main
import ("fmt")
type rectangle struct{
	length, height float64
}

func (r rectangle)Area()float64{
	return r.length * r.height
}

func main(){
	rect1 := rectangle{1.33, 22.3}
	fmt.Println(rect1.Area())
}