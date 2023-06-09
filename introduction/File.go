package main
import(
	"fmt"
	"os"
	"log"
	"strconv"
	"bufio"
)
var pl = fmt.Println

func main(){
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArry := []int{2,3,5,7}
	var sPrimeArry []string
	for _, vl :=range iPrimeArry{
		sPrimeArry = append(sPrimeArry, strconv.Itoa(vl))
	}

	for _, vl := range sPrimeArry{
		_, err := f.WriteString(vl + "\n")
		if err != nil{
			log.Fatal(err)
		}
	}

	f, err = os.Open("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	scan1 := bufio.NewScanner(f)
	for scan1.Scan(){
		pl("Prime numbers are", scan1.Text())
	}

	if err := scan1.Err(); err != nil{
		log.Fatal(err)
	}


}