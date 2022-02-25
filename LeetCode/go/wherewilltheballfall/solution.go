package wherewilltheballfall

func getGridValue(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	return grid[i][j]
}

func wheretheballfall(grid [][]int, i, j int) int {
	if i >= len(grid) {
		return j
	}
	currentValue := getGridValue(grid, i, j)
	if currentValue == 1 && getGridValue(grid, i, j+1) == 1 {
		return wheretheballfall(grid, i+1, j+1)
	} else if currentValue == -1 && getGridValue(grid, i, j-1) == -1 {
		return wheretheballfall(grid, i+1, j-1)
	}
	return -1
}

func findBall(grid [][]int) []int {
	result := make([]int, len(grid[0]))
	for j := range result {
		result[j] = wheretheballfall(grid, 0, j)
	}
	return result
}
