package main

import (
	"fmt"
	"headFirstGo/chapter5/array"
	"headFirstGo/chapter5/readfile"
	"log"
)

func main() {
	array.TestArray()
	avg := array.Average([3]float64{71.8, 56.2, 89.5})
	fmt.Println(avg)

	numbers, err := readfile.GetFloats("/Users/jesse/work/project/headFirstGo/chapter5/data/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64
	for _, number := range numbers {
		sum = +number
	}
	fmt.Println(numbers)
	sampleCount := float64(len(numbers))
	fmt.Printf("Average:%0.2f\n", sum/sampleCount)
}
