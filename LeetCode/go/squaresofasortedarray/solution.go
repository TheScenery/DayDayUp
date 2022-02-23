package squaresofasortedarray

import "sort"

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = num * num
	}
	sort.Ints(result)
	return result
}

func sortedSquares1(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	r := length - 1
	i := 0
	j := r
	for {
		if i > j {
			break
		}
		x := nums[i] * nums[i]
		y := nums[j] * nums[j]
		if x > y {
			result[r] = x
			i++
		} else {
			result[r] = y
			j--
		}
		r--

	}
	return result
}
