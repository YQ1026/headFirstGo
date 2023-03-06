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
