package platesbetweencandles

func findLeft(candles []int, key int) int {
	for _, c := range candles {
		if c >= key {
			return c
		}
	}
	return -1
}

func findRight(candles []int, key int) int {
	n := len(candles)
	for i := n - 1; i >= 0; i-- {
		if candles[i] <= key {
			return candles[i]
		}
	}
	return -1
}

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	candles := make([]int, 0)
	prePlates := make([]int, n)
	for i, ch := range s {
		if ch == '|' {
			candles = append(candles, i)
		}
		if i == 0 {
			if ch == '*' {
				prePlates[i] = 1
			}
			continue
		}
		if ch == '|' {
			prePlates[i] = prePlates[i-1]
		} else {
			prePlates[i] = prePlates[i-1] + 1
		}
	}
	m := len(queries)
	result := make([]int, m)
	for i, q := range queries {
		left := findLeft(candles, q[0])
		right := findRight(candles, q[1])
		if left == -1 || right == -1 || left >= right {
			result[i] = 0
		} else {
			result[i] = prePlates[right] - prePlates[left]
		}
	}
	return result
}
