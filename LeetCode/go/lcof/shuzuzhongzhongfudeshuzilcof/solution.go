package shuzuzhongzhongfudeshuzilcof

func findRepeatNumber(nums []int) int {
	hasTable := make(map[int]bool)
	for _, n := range nums {
		if hasTable[n] {
			return n
		}
		hasTable[n] = true
	}
	return -1
}
