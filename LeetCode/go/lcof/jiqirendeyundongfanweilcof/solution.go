package jiqirendeyundongfanweilcof

type Position struct {
	i int
	j int
}

func isValidPosition(m int, n int, k int, i, j int, visitedMap map[Position]bool) bool {
	if i < 0 || i >= m || j < 0 || j >= n {
		return false
	}
	if visitedMap[Position{i, j}] {
		return false
	}
	sum := 0
	for i > 0 {
		sum = sum + i%10
		i = i / 10
	}
	for j > 0 {
		sum = sum + j%10
		j = j / 10
	}

	return sum <= k
}

func visit(m int, n int, k int, i, j int, visitedMap map[Position]bool) int {
	count := 0
	if !isValidPosition(m, n, k, i, j, visitedMap) {
		return 0
	}
	count++
	visitedMap[Position{i, j}] = true
	count += visit(m, n, k, i-1, j, visitedMap)
	count += visit(m, n, k, i+1, j, visitedMap)
	count += visit(m, n, k, i, j-1, visitedMap)
	count += visit(m, n, k, i, j+1, visitedMap)
	return count
}

func movingCount(m int, n int, k int) int {
	return visit(m, n, k, 0, 0, map[Position]bool{})
}
