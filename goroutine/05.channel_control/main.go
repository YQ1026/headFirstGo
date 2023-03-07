package main

import (
	"fmt"
	"time"
)

func Process(ch chan int)  {
	time.Sleep(time.Second)
	ch <-1
}

func main(){
	// 创建一个包含10个元素的切片，类型为channel。
	var channels []chan int
	channels = make([]chan int,10)
	for i:=0;i<10;i++{
		channels[i]=make(chan int)
		go Process(channels[i])
	}

	for i,ch := range channels{
		fmt.Println(i,<-ch)
		fmt.Printf("Routine %v quit!\n",i)
	}
}
