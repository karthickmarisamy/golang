package main
import (
	"fmt"
	"time"
	"math/rand"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)
var pl =fmt.Println
func main(){
	randSec := time.Now().Unix()
	rand.Seed(randSec)
	randNum := rand.Intn(50)+1;
	for true{
	pl("Guess the numner between 0 and 50:\n")
	pl("Random number is", randNum)
	pl("Your guessing number:\n")
	inputRead := bufio.NewReader(os.Stdin)
	finalInput, err := inputRead.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	finalInput = strings.TrimSpace(finalInput)
	iGuess, err := strconv.Atoi(finalInput)
	if err != nil{
		log.Fatal(err)
	}

	if iGuess > randNum{
		pl("Your gussed number is greater than the number")
	}else if iGuess < randNum{
		pl("Your guessed number is lower than the number")
	}else{
		pl("You guessed it")
		break
	}
}

}