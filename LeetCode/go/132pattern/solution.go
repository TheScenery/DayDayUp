package pattern132

import "math"

func find132pattern(nums []int) bool {
	m := len(nums)
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			for k := j + 1; k < m; k++ {
				if nums[i] < nums[k] && nums[k] < nums[j] {
					return true
				}
			}
		}
	}
	return false
}

func find132pattern1(nums []int) bool {
	m := len(nums)
	stack := make([]int, 0)
	maxK := math.MinInt
	for i := m - 1; i >= 0; i-- {
		n := nums[i]
		if n < maxK {
			return true
		}
		if len(stack) == 0 {
			stack = append(stack, n)
		} else if n < stack[len(stack)-1] {
			stack = append(stack, n)
		} else {
			for len(stack) > 0 && n > stack[len(stack)-1] {
				if stack[len(stack)-1] > maxK {
					maxK = stack[len(stack)-1]
				}
				stack = stack[0 : len(stack)-1]
			}
			stack = append(stack, n)
		}
	}
	return false
}
