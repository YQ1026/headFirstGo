package inRange

func InRange(min int, max int, numbers ...int) []int {

	var saveRangeValue []int
	for _, number := range numbers {
		if number > min && number < max {
			saveRangeValue = append(saveRangeValue, number)
		}
	}
	return saveRangeValue
}
