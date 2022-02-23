package matrix01

import "math"

type Point struct {
	Row int
	Col int
}

//BFS
func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	result := make([][]int, m)
	visitedMap := make(map[Point]bool)
	queue := make([]Point, 0)
	for i := range mat {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
				point := Point{Row: i, Col: j}
				visitedMap[point] = true
				queue = append(queue, point)
			}
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			nextPoint := Point{Row: currentPoint.Row + dir[0], Col: currentPoint.Col + dir[1]}
			if !visitedMap[nextPoint] && nextPoint.Row >= 0 && nextPoint.Row < m && nextPoint.Col >= 0 && nextPoint.Col < n {
				result[nextPoint.Row][nextPoint.Col] = 1 + result[currentPoint.Row][currentPoint.Col]
				visitedMap[nextPoint] = true
				queue = append(queue, nextPoint)
			}
		}
	}
	return result
}

//DP
func updateMatrix1(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	maxSteps := m + n
	result := make([][]int, m)
	for i := range mat {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
			} else {
				result[i][j] = maxSteps
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				if i-1 >= 0 {
					result[i][j] = int(math.Min(float64(result[i][j]), float64(result[i-1][j]+1)))
				}
				if j-1 >= 0 {
					result[i][j] = int(math.Min(float64(result[i][j]), float64(result[i][j-1]+1)))
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				result[i][j] = int(math.Min(float64(result[i][j]), float64(result[i+1][j]+1)))
			}
			if j+1 < n {
				result[i][j] = int(math.Min(float64(result[i][j]), float64(result[i][j+1]+1)))
			}
		}
	}
	return result
}
