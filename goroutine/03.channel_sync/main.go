package main

import "fmt"

func abc(channel chan string){
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func xyz(channel chan string){
	channel <- "x"
	channel <- "y"
	channel <- "z"
}

func main(){
	chan1 := make(chan string)
	chan2 := make(chan string)

	go abc(chan1)
	go xyz(chan2)

	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
}