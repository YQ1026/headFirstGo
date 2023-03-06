package greeting

import "fmt"

// 首字母大写可以导出
func Hello() {
	fmt.Println("Hello")
}

func Hi() {
	fmt.Println("Hi")
}
