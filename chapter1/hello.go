package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var originalCount int = 10
var eatenCount int = 4

func main() {
	// fmt.Println(reflect.TypeOf("str"))
	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(replacer.Replace(broken))

	// 设置从键盘获取文本的缓冲读取器
	fmt.Println("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)

	// 返回用户输入的所有内容 直到按下enter键为止
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(input)

	fileInfo, err1 := os.Stat("/Users/jesse/work/project/src/headFirstGo/chapter01/a.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(fileInfo.Size())

}
