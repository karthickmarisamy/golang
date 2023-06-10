package main
import (
	"fmt"
	"os"
	"log"
	"errors"
)
var pl = fmt.Println
func main(){
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist){
		pl("File not found")
	}else{
	  f, err :=os.OpenFile("data.txt", 
	  				os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0644)
		if err != nil{
			log.Fatal(err)
		}
		
		defer f.Close()

		if _, err := f.WriteString("12\n"); err !=nil {
			log.Fatal(err)
		}
	}
}