package main

import (
	"fmt"
	"headFirstGo/chapter6/autoFunc"
	"headFirstGo/chapter6/example"
	"headFirstGo/chapter6/inRange"
	"headFirstGo/chapter6/maximum"
	"log"
	"os"
	"strconv"
)

func main() {
	example.SliceExample()

	// 从参数获取输入浮点数的平均数
	arguments := os.Args[1:]
	var sum float64 = 0.0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}

	sampleCount := float64(len(arguments))
	fmt.Printf("Average:%0.2f\n", sum/sampleCount)

	autoFunc.SeveralInts(1, 2, 3, 4, 5, 6, 67, 7, 78, 8, 8, 8, 8, 8, 8, 8)

	// 获取最大数字
	maxFloat := maximum.Maximum(1.1, 2.3, 9.9, 23.1, 456.9)
	fmt.Println(maxFloat)

	// 获取范围内的数值
	inRangeResult := inRange.InRange(1, 900, 999, 200, 399, 98766)
	fmt.Println(inRangeResult)
}
