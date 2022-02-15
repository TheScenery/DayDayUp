package rotatearray

func rotate(nums []int, k int) {
	l := len(nums)
	result := make([]int, l)
	normalizedK := k % l
	for i := 0; i < l; i++ {
		sourceIndex := i - normalizedK
		if sourceIndex < 0 {
			sourceIndex = l + sourceIndex
		}
		result[i] = nums[sourceIndex]
	}
	for i := 0; i < l; i++ {
		nums[i] = result[i]
	}
}

func reverse(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate1(nums []int, k int) {
	l := len(nums)
	normalizedK := k % l
	reverse(nums, 0, l-1)
	reverse(nums, 0, normalizedK-1)
	reverse(nums, normalizedK, l-1)
}
