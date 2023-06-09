package main
import ("fmt")
var pl = fmt.Println

func calculateAvg(paramArry ... float64) float64{
	var sum float64 = 0.0
	var numSize float64  = float64(len(paramArry))
	for _, vl :=range paramArry{
		sum += vl;
	}
	return (sum / numSize)
}

func main(){
	var floatVal = []float64{2.33, 22.3, 1.2, 3.4}
	fmt.Printf("The Average value is : %.3f ", calculateAvg(floatVal ...))
}