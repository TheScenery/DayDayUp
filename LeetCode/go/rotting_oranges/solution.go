package rottingoranges

type Point struct {
	Row int
	Col int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	freshCount := 0
	currentRottingOranges := make([]Point, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				point := Point{Row: i, Col: j}
				currentRottingOranges = append(currentRottingOranges, point)
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}
	if freshCount == 0 {
		return 0
	}

	minutes := -1
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	nextMinuteRottingOranges := make([]Point, 0)
	for len(currentRottingOranges) > 0 {
		orange := currentRottingOranges[0]
		currentRottingOranges = currentRottingOranges[1:]
		for _, dir := range dirs {
			next := Point{Row: orange.Row + dir[0], Col: orange.Col + dir[1]}
			if next.Row >= 0 && next.Row < m && next.Col >= 0 && next.Col < n && grid[next.Row][next.Col] == 1 {
				nextMinuteRottingOranges = append(nextMinuteRottingOranges, next)
				freshCount--
				grid[next.Row][next.Col] = 0
			}
		}
		if len(currentRottingOranges) == 0 {
			currentRottingOranges = nextMinuteRottingOranges
			nextMinuteRottingOranges = make([]Point, 0)
			minutes++
		}
	}
	if freshCount > 0 {
		return -1
	}
	return minutes
}
