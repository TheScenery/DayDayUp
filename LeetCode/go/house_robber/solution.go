package houserobber

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	p, q := nums[0], max(nums[0], nums[1])
	r := 0
	for i := 3; i <= len(nums); i++ {
		r = max(nums[i-1]+p, q)
		p, q = q, r
	}
	return r
}
