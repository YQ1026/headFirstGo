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

	count := 10

	for {
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

		if count < 1 {
			fmt.Println("对不起 你没有猜对  它是:", targetNumber)
			break
		} else {
			if guessNum > targetNumber {
				fmt.Println("哎呀 你猜高了")
				count--
				fmt.Println("当前剩余猜测次数:", count)
				fmt.Println(strings.Repeat("-", 100))
			} else if guessNum < targetNumber {
				fmt.Println("哎呀 你猜低了")
				count--
				fmt.Println("当前剩余猜测次数:", count)
				fmt.Println(strings.Repeat("-", 100))
			} else {
				fmt.Println("你猜对了")
				break
			}
		}

	}

}
