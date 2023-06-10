package main
import (
	"fmt"
	"os"
	"log"
	"bufio"
)
var pl = fmt.Println
func main(){
	f, err :=os.OpenFile("data1.txt", 
			os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644 )
	if err != nil{
		log.Fatal(err)
	}
	
	defer f.Close()
	pl("The file is ready for writing your input")

	cString := bufio.NewReader(os.Stdin)
	cStr, err := cString.ReadString('\n')

	if _, err := f.WriteString(cStr); err != nil {
		log.Fatal(err)
	}

	pl("The file is written successfully")

	pl("The current file contents are as follows")

	f, err = os.Open("data1.txt")
	rFile := bufio.NewScanner(f)

	for rFile.Scan(){
		pl(rFile.Text())
	}
}