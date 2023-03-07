package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue(float64) string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("我是方法，但是没有参数。")
}

func (m MyType) MethodWithParameters(f float64){
	fmt.Printf("我是方法，会传入一个参数:%v,接收的对象值是:%v\n",f,m)
}

func (m MyType) MethodWithReturnValue(f float64) string {
	return fmt.Sprintf("我是一个带有传入参数的并且有返回值的方法%v,接收的对象值是:%v\n",f,m)
}

func (my MyType) MethodNotInInterface()  {
	fmt.Println("我是MyType方法的对象，但是我不属于接口，我不能被接口调用，但是可以被对象调用")
}
