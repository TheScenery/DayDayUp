package knightprobabilityinchessboard

import "math"

type position struct {
	Row int
	Col int
}

func isIn(n int, p position) bool {
	return p.Row >= 0 && p.Row < n && p.Col >= 0 && p.Col < n
}

func move(n int, start position) []position {
	result := make([]position, 0)
	possibleMove := []position{{Row: 2, Col: 1}, {Row: 2, Col: -1}, {Row: -2, Col: 1}, {Row: -2, Col: -1}, {Row: 1, Col: 2}, {Row: -1, Col: 2}, {Row: 1, Col: -2}, {Row: -1, Col: -2}}
	for _, v := range possibleMove {
		nextPosition := position{Row: start.Row + v.Row, Col: start.Col + v.Col}
		if isIn(n, nextPosition) {
			result = append(result, nextPosition)
		}
	}
	return result
}

// 内存爆炸
func knightProbability1(n int, k int, row int, column int) float64 {
	totalCount := math.Pow(8, float64(k))
	currentPositions := []position{{Row: row, Col: column}}
	for k > 0 {
		nextPositions := make([]position, 0)
		for _, currentPosition := range currentPositions {
			nextPositions = append(nextPositions, move(n, currentPosition)...)
		}
		currentPositions = nextPositions
		k--
	}
	return float64(len(currentPositions)) / totalCount
}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, k+1)
	for index := 0; index <= k; index++ {
		dp[index] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[index][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if index == 0 {
					dp[index][i][j] = 1
				} else {
					result := move(n, position{Row: i, Col: j})
					for _, r := range result {
						dp[index][i][j] = dp[index][i][j] + dp[index-1][r.Row][r.Col]/8
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
