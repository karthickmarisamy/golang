package stuff

import (
	"fmt"
	"strconv"
)

var Name := "Karthick"

func IntArrytoStrArry(iArry = []int)[]string{
	sArry := []string{}
	for _, i := range iArry{
		val := strconv.Itoa(i)
		sArry.append(sArry, val)
	}
	return sArry;
}