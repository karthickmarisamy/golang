package main
import(
	"fmt"
)
func main(){
	fmt.Printf("%d %s %c %t %o %x\n", 123, "Test", "a", true, 1, 1 )
	fmt.Printf("%9f\n", 23.4)
	fmt.Printf("%.2f\n", 23.4234323)
	fmt.Printf("%9.f\n", 23.4234323)	
}