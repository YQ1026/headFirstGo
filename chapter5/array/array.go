package array

import (
	"fmt"
	"time"
)

func TestArray() {
	var dates [3]time.Time
	dates[0] = time.Now()
	dates[1] = time.Now()
	dates[2] = time.Now()

	//fmt.Printf("%#v\n", dates)
	for i := 0; i < len(dates); i++ {
		fmt.Println(dates[i])
	}

	for index, value := range dates {
		fmt.Printf("index:%v,value:%v\n", index, value)
	}

	// 如果不需要index 或者value 可以使用_代替:
	for _, value := range dates {
		fmt.Printf("value:%v\n", value)
	}
}

func Average(nums [3]float64) float64 {
	var sum float64
	for _, value := range nums {
		sum = +value
	}
	avg := sum / float64(len(nums))
	return avg
}
