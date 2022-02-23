package maxareaofisland

func processIsland(grid [][]int, m, n, r, c int) int {
	if r < 0 || r >= m || c < 0 || c >= n {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	area := 1
	grid[r][c] = 0
	area += processIsland(grid, m, n, r-1, c)
	area += processIsland(grid, m, n, r, c-1)
	area += processIsland(grid, m, n, r+1, c)
	area += processIsland(grid, m, n, r, c+1)
	return area
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxArea := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				area := processIsland(grid, m, n, r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
