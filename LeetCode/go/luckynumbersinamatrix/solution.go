package luckynumbersinamatrix

func luckyNumbers(matrix [][]int) []int {
	maxValueMap := make(map[int]int)
	minValueMap := make(map[int]int)
	resultMap := make(map[int]bool)
	for i, x := range matrix {
		for j, y := range x {
			minY, ok := minValueMap[i]
			if !ok || y < minY {
				minValueMap[i] = y
			}
			maxX, ok := maxValueMap[j]
			if !ok || y > maxX {
				maxValueMap[j] = y
			}
		}
	}

	result := make([]int, 0)
	for _, value := range minValueMap {
		resultMap[value] = true
	}
	for _, value := range maxValueMap {
		if resultMap[value] {
			result = append(result, value)
		}
	}
	return result
}
