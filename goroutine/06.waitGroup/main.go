package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func(name string) {
		fmt.Printf("goroutine %v done\n", name)
		wg.Done()
	}("goroutine1")

	go func(name1 string) {
		fmt.Printf("goroutine %v done\n", name1)
		wg.Done()
	}("goroutine2")

	wg.Wait()
}
