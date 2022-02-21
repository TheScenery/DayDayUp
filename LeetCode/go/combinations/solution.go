package combinations

func combine(n int, k int) [][]int {
	if n < k || n <= 0 || k < 0 {
		return [][]int{}
	}
	if k == 1 {
		result := make([][]int, n)
		for i := 0; i < n; i++ {
			result[i] = []int{i + 1}
		}
		return result
	}
	results := make([][]int, 0)
	for i := n; i >= k; i-- {
		result := combine(i-1, k-1)
		for _, r := range result {
			results = append(results, append(r, i))
		}
	}
	return results
}

func combine1(n int, k int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, 0)
	var backtrack func(int)
	backtrack = func(start int) {
		if len(temp)+(n-start+1) < k {
			return
		}
		if len(temp) == k {
			com := make([]int, k)
			copy(com, temp)
			ans = append(ans, com)
			return
		}

		temp = append(temp, start)
		backtrack(start + 1)
		temp = temp[0 : len(temp)-1]

		backtrack(start + 1)

	}
	backtrack(1)
	return ans
}
