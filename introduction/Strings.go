package main
import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main(){
	pl("String Length", len("Hello this is go language"))
	pl("String contains Hello", strings.Contains("Hello this is string", "Hello"))
	pl("String replace with Hi", strings.Replace("Hello everyone", "Hello", "Hi", -1))
	pl("String prefix with Hello", strings.HasPrefix("Hello All", "Hello"))
	pl("String suffix with Hello", strings.HasSuffix("Hello golang", "golang"))
	pl("Remove space in the string", strings.TrimSpace("\nHello Everyone\n"))
	pl("String split", strings.Split("a-b-c", "-"))
	pl("String to convert into uppercase", strings.ToUpper("hello all"))
	pl("String to convert into lowercase", strings.ToLower("HELLO ALL"))
	pl("Get string index", strings.Index("Hello", "o"))
}