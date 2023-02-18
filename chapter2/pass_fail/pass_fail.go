package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade:")
	r := bufio.NewReader(os.Stdin)   // 读取键盘输入
	input, err := r.ReadString('\n') // 读取用户输入的内容 直到他们按了enter键盘
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input) // 删除input输入前后的换行符(空格等)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "good"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)

	/*
		正常情况下 同一变量名在同一作用域声明两次 会出现编译错误,
		但是只要短变量声明中 至少有一个变量名是新的,是被允许的。
		新变量名视为声明,老的视为赋值。
	*/

	a := 1
	a, b := 2, 3
	fmt.Println(a, b)
}
