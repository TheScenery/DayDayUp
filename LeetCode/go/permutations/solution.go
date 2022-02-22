package permutations

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, 0)
	visitedMap := make(map[int]bool)
	n := len(nums)

	var backtrack func()
	backtrack = func() {
		if len(temp) == n {
			per := make([]int, len(temp))
			copy(per, temp)
			ans = append(ans, per)
			return
		}

		for i := 0; i < n; i++ {
			cur := nums[i]
			if !visitedMap[cur] {
				temp = append(temp, cur)
				visitedMap[cur] = true
				backtrack()
				temp = temp[:len(temp)-1]
				visitedMap[cur] = false

			}
		}
	}
	backtrack()
	return ans
}
