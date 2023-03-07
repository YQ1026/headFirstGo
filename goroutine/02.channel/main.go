package main

import "fmt"

//
//var myChannel chan float64 // 声明一个变量来保存channel
//myChannel = make(chan float64) // 实际创建channel
//
//myChannel := make(chan float64) // 创建一个channel并立即声明一个变量

// 将channel作为一个参数
func greeting(myChannel chan string){
	myChannel <- "hello chan!" // 通过channel发送一个值
}


func main(){
	myChannel := make(chan string)
	go greeting(myChannel) //将channel传递给在新goroutine中运行的函数

	receiver := <- myChannel // 从channel中接收值。
	fmt.Println(receiver)
}