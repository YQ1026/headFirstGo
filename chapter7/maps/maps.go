package maps

import (
	"fmt"
)

func GetCounters() {
	counters := map[string]int{"a": 3, "b": 4}

	valueB, ok := counters["b"]
	if ok {
		fmt.Println(valueB)
	}
	valueC, ok := counters["c"]
	if ok {
		fmt.Println(valueC)
	}
}

func GetMapOut() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
		fmt.Println(counts[item])
	}
	fmt.Println(counts)

	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
}
