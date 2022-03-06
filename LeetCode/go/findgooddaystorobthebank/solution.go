package findgooddaystorobthebank

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	result := make([]int, 0)
	if n < 2*time+1 {
		return result
	}
	dpNonIncreasing := make([]int, n)
	dpNonIncreasing[0] = 0
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			dpNonIncreasing[i] = dpNonIncreasing[i-1] + 1
		} else {
			dpNonIncreasing[i] = 0
		}
	}
	dpNonDecreasing := make([]int, n)
	dpNonDecreasing[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			dpNonDecreasing[i] = dpNonDecreasing[i+1] + 1
		} else {
			dpNonDecreasing[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		if dpNonIncreasing[i] >= time && dpNonDecreasing[i] >= time {
			result = append(result, i)
		}
	}
	return result
}
