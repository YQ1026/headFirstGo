package main

import (
	"fmt"
	"log"
)

type FakeError string

func (f FakeError) Error() string  {
	return string(f)
}


type OverheatError float64

func (o OverheatError) Error() string{
	return fmt.Sprintf("%0.2f",o)
}

func checkTemperature(actual,safe float64) error {
	excess := actual - safe
	if excess >0 {
		return OverheatError(excess)
	}
	return nil
}


func main(){
	//var err error
	//err = FakeError("我是传入的错误值")
	//fmt.Println(err)

	var err error = checkTemperature(130.122,100.0)
	if err != nil{
		log.Fatal(err)
	}
}