package main

import (
	"fmt"
	"headFirstGo/interface/02.mypkg/mypkg"
	_ "headFirstGo/interface/02.mypkg/mypkg"
)

func main(){
	var value mypkg.MyInterface // 声明一个接口类型的变量
	value = mypkg.MyType(1) //实例化


	value.MethodWithoutParameters()
	value.MethodWithParameters(1.02)
	ret:= value.MethodWithReturnValue(1.01)
	fmt.Println(ret)
}
