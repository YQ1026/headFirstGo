package main

import "fmt"

type Switch string

func (s *Switch) Toggle(){
	if *s == "on" {
		*s = "off"
	}else {
		*s = "on"
	}
	fmt.Println("*s=",*s)
}


type Toggleable interface {
	Toggle()
}

func main()  {
	s := Switch("off") //这里接收的是对象实例化的传参
	var t Toggleable = &s
	t.Toggle()
	t.Toggle()
}
