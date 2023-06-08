package main
import (
	"fmt"
)
func main()  {
	var arr1 = [5]int{}
	arr1[0] = 2
	var arr2 = [5]int{1,3,4,5,6}
	
	for i := 0; i < len(arr1); i++{
		fmt.Printf("%d : %d\n", i, arr1[i])
	}

	for i:=0; i < len(arr2); i++ {
		fmt.Printf("%d : %d\n", i, arr2[i])
	}

	var arr3 = [2][2]int{
		{1,2},
		{3,4},
	}

	for i := 0; i < len(arr3); i++{
		for j:=0; j < len(arr3); j++{
			fmt.Printf("%d\n", arr3[i][j])
		}
	}
}