package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	targetNumber := rand.Intn(100) + 1 //[0-100) 包含0 但不包含100

	fmt.Println("我已经选了一个1-100之间的随机数,小朋友 快来猜猜它是几?")
	fmt.Println(strings.Repeat("-", 100))

	success := false

	for guessCount := 0; guessCount < 10; guessCount++ {
		r := bufio.NewReader(os.Stdin)
		input, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guessNum, err := strconv.Atoi(input) // 将输入的字符串转换为整数
		if err != nil {
			log.Fatal(err)
		}

		if guessNum > targetNumber {
			fmt.Println("哎呀 你猜高了")
			fmt.Printf("当前已猜: %d 次\n", guessCount+1)
			fmt.Println(strings.Repeat("-", 100))
		} else if guessNum < targetNumber {
			fmt.Println("哎呀 你猜低了")
			fmt.Printf("当前已猜: %d 次\n", guessCount+1)
			fmt.Println(strings.Repeat("-", 100))
		} else {
			success = true
			fmt.Println("你猜对了")
			break
		}
	}
	if !success {
		fmt.Printf("你没有猜对我的数字,实际上它是:%d", targetNumber)
	}
}
