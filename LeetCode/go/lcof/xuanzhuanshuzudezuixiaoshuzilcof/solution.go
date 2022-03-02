package xuanzhuanshuzudezuixiaoshuzilcof

func minArray(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		if i > 0 && numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}

func minArray1(numbers []int) int {
	n := len(numbers)
	low := 0
	high := n - 1
	for low < high {
		p := low + (high-low)/2
		if numbers[p] < numbers[high] {
			high = p
		} else if numbers[p] > numbers[high] {
			low = p + 1
		} else {
			high = high - 1
		}
	}
	return numbers[low]
}
