package main

import "fmt"

func reportNap(name string,delay int){
	for i:=0;i<delay;i++{
		fmt.Println(name,"sleeping")
	}
	fmt.Println(name,"wake up.")
}

func send(myChannel chan string){
	reportNap("sending goroutine",2)
	fmt.Println("---sending value a---")
	myChannel <-"a"
	fmt.Println("---sending value b---")
	myChannel <- "b"
}

func main(){
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("main channel receiver",5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

}