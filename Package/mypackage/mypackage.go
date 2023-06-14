package stuff

import (
	"strconv"
)

var Name string = "Karthick"

func IntArrytoStrArry(iArry []int)[]string{
	var sArry []string
	for _, i := range iArry{
		sArry = append(sArry, strconv.Itoa(i))
	}
	return sArry;
}