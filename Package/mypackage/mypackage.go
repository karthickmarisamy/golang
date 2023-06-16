package stuff

import (
	"strconv"
	"errors"
	"time"
)

var Name string = "Karthick"

func IntArrytoStrArry(iArry []int)[]string{
	var sArry []string
	for _, i := range iArry{
		sArry = append(sArry, strconv.Itoa(i))
	}
	return sArry;
}


type Date struct{
	day int
	month int
	year int
}

func (d *Date) SetDay(day int)error{
	if(day < 1 && day > 31){
		errors.New("Invalid Day")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(month int)error{
	if(month < 1 && month > 12){
		errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

func (d *Date) SetYear(year int)error{
	if(year < 1875 && year > time.Now().Year()){
		errors.New("Invalid Year")
	}
	d.year = year
	return nil
}

func (d *Date) Date()int{
	return d.day
}

func (d *Date) Month()int{
	return d.month
}

func (d *Date) Year()int{
	return d.year
}

