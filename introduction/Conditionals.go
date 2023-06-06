package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	iAge :=17
	if (iAge >= 1) && (iAge <=18){
		pl("You are minor")
	}else if(iAge >19){
		pl("You are Major")
	}else{
		pl("Infant")
	}
}