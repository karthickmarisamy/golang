package main
import (
	"fmt"
	"time"
)

var pl = fmt.Println
func main(){
	timeData := time.Now()
	pl(timeData.Year(), timeData.Month(), timeData.Day())
	pl(timeData.Hour(), timeData.Minute(), timeData.Second())
}