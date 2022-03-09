package juzhenzhongdelujinglcof

type Point struct {
	i int
	j int
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	possibleStart := make([]Point, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				possibleStart = append(possibleStart, Point{i, j})
			}
		}
	}

	currentWord := ""
	visited := make(map[Point]bool)
	dirs := []Point{{i: -1, j: 0}, {i: 1, j: 0}, {i: 0, j: -1}, {i: 0, j: 1}}
	var backTracking func(int, int) bool
	backTracking = func(i, j int) bool {
		if currentWord == word {
			return true
		}
		if currentWord != word[0:len(currentWord)] {
			return false
		}

		for _, dir := range dirs {
			newPosition := Point{i: i + dir.i, j: j + dir.j}
			if newPosition.i < 0 || newPosition.i >= m || newPosition.j < 0 || newPosition.j >= n {
				continue
			}
			if !visited[newPosition] {
				currentWord += string(board[newPosition.i][newPosition.j])
				visited[newPosition] = true
				r := backTracking(newPosition.i, newPosition.j)
				if r {
					return true
				}

				currentWord = currentWord[0 : len(currentWord)-1]
				visited[newPosition] = false

			}
		}

		return false
	}

	for _, start := range possibleStart {
		currentWord = string(board[start.i][start.j])
		visited = map[Point]bool{start: true}
		r := backTracking(start.i, start.j)
		if r {
			return true
		}
	}
	return false
}
