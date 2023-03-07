package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("我发出的是口哨声。")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("我发出的是号角声。")
}

// 定义接口 代表任何含有MakeSound的方法
type NoiseMaker interface {
	MakeSound()
}

// 定义函数参数为interface类型。
func play(n NoiseMaker) {
	n.MakeSound()
}

// 普通调用方法
func call() {
	var n NoiseMaker
	n = Whistle("口哨")
	n.MakeSound()

	n = Horn("号角")
	n.MakeSound()
}

// 新建对象robot 满足MakeSound方法基础上额外加一个walk方法
type Robot string

func (r Robot) MakeSound()  {
	fmt.Println("我是机器人 满足MakeSound方法。")
}

func (r Robot) Walk()  {
	fmt.Println("我是walk方法，但是不属于any interface.")
}

func main() {
	//函数式的调用方法
	play(Horn("号角"))
	play(Whistle("口哨声"))

	// 普通调用方法
	call()

	var i NoiseMaker
	i = Robot("我是机器人")
	i.MakeSound()

	// i.Walk undefined (type NoiseMaker has no field or method Walk)
	// NoiseMaker 只能调用接口定义好的方法。
	// i.Walk()

}
