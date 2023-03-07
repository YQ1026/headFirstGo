package main

import "fmt"

type Animal interface{
	Eat()
}

type Dog string

func (d Dog) Eat(){
	fmt.Println(d,"wang wang!")
}

func (d Dog) Run(){
	fmt.Println(d,"speed up!")
}

func main(){
	var a Animal = Dog("旺财")
	a.Eat()

	// interface里面我们只定义了Eat()方法，Dog对象里我们定义了Eat()/Run()方法，此时 我们直接调用Run方法，报错如下:
	// a.Run undefined (type Animal has no field or method Run)
	// a.Run()

	// 这里使用断言从接口里取回Dog对象的具体类型
	var d Dog = a.(Dog)
	d.Run()


}


