package main

import (
	"fmt"
	"headFirstGo/chapter8/magazine"
)

func main() {

	employee := &magazine.Employee{
		Name:   "Jesse",
		Salary: 10000.1,
		HomeAddress: magazine.Address{
			City:       "北京",
			PostalCode: 10000,
		},
	}

	fmt.Println(employee)
}
