package main

import (
	"fmt"
	 "headFirstGo/interface/01.gadget/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player,songs []string){
	for _,song := range songs{
		device.Play(song)
	}
	device.Stop()
}

// Run函数
func OperatorForBoth(){
	var player Player

	// 播放歌曲
	player = gadget.TapePlayer{}
	songs :=[]string{"歌曲1","歌曲2","歌曲3"}
	playList(player,songs)

	// 录制歌曲
	player = gadget.TapeRecorder{}
	playList(player,songs)
}

// 测试断言失败
func TryOut(player Player){
	player.Play("测试1")
	player.Stop()

	// 断言类型是TapeRecorder
	recorder,ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}else{
		fmt.Println("TapePlayer was not a TapeRecorder")
	}

}



// 测试类型断言
type Human interface {
	Eat()
}

type People struct {
	Name string
	Age int
}

func (p People) Eat(){
	fmt.Printf("我是%v,我今年%v岁了，我是吃方法。",p.Name,p.Age)
}

func (p People) Say(){
	fmt.Printf("我是%v,我今年%v岁了，我是说方法。",p.Name,p.Age)
}

func CallMethod(){
	var h Human = People{"Jesse",28}
	h.Eat()

	//使用类型断言取回具体类型
	var people People = h.(People)

	// 调用具体类型上的(而不是接口上的)的定义
	people.Say()
}

func main()  {
	// CallMethod()
	// OperatorForBoth()
	TryOut(gadget.TapeRecorder{})

	// 断言类型是TapeRecorder 实际上是TapePlayer,panic。
	TryOut(gadget.TapePlayer{})
}
