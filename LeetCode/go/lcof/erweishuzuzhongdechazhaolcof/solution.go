package erweishuzuzhongdechazhaolcof

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	n := 0
	if m != 0 {
		n = len(matrix[0])
	}
	for i := 0; i < m; i++ {
		s2 := 0
		e2 := n - 1
		for s2 <= e2 {
			m2 := s2 + (e2-s2)/2
			if matrix[i][m2] == target {
				return true
			} else if matrix[i][m2] < target {
				s2 = m2 + 1
			} else {
				e2 = m2 - 1
			}
		}
	}

	return false
}
