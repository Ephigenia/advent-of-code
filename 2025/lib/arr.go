package lib

// todo move to lib
func ArrMaxAndIndex(vals []int) (int, int) {
	max := -1
	index := -1
	for i, digit := range vals {
		if digit > max {
			max = digit
			index = i
		}
	}
	return max, index
}
