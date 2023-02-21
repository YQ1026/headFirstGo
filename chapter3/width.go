package main

import (
	"fmt"
	"log"
	"strings"
)

func paintNeeded(witdth float64, height float64) (float64, error) {
	if witdth < 0 {
		return 0, fmt.Errorf("A width of %0.2f is invalid", witdth)
	}

	if height < 0 {
		return 0, fmt.Errorf("A height of %0.2f is invalid", height)
	}
	area := witdth * height
	return area / 10.0, nil
}

func pointerTest() {
	myInt := 4
	fmt.Println(myInt)

	myIntPointer := &myInt
	*myIntPointer = 88
	fmt.Printf("pointer:%v,value: %v\n", myIntPointer, *myIntPointer)
}

func pointerTest2() {
	var myInt int
	myIntPoint := &myInt

	*myIntPoint = 99
	fmt.Println(myInt, *myIntPoint)
}

func pointerTest3(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	pointerTest2()
	pointerTest()
	truth := true
	pointerTest3(&truth)
	fmt.Println(truth)

	fmt.Printf("%12s|%s\n", "Product", "Cost in cents")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("%12s|%3d\n", "Stamps", 50)
	fmt.Printf("%12s|%3d\n", "Paper Clips", 5)
	fmt.Printf("%12s|%3d\n", "Tape", 99)

	//格式化小数宽度测试
	fmt.Printf("%%7.3f: %7.3f\n", 3.1415926) // 7.3表示最小宽度,小数点后面的表示四舍五入到几位。
	fmt.Printf("%%.1f: %.1f\n", 3.1415926)   // .1f表示四舍五入到1位 无填充。

	// paintNeeded
	amount, err := paintNeeded(1.1, -3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("0.2f liters needed!\n", amount)

}
