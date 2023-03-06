package main

import (
	"fmt"
	"headFirstGo/chapter4/greeting"
)
import "headFirstGo/chapter4/calc"

func main() {
	greeting.Hi()
	greeting.Hello()
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(22, 11))
}
