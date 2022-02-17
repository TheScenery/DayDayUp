package floodfill

func fill(image [][]int, m, n, sr, sc, oldColor, newColor int) {
	image[sr][sc] = newColor
	nextSr := sr - 1
	if 0 <= nextSr && nextSr < m && image[nextSr][sc] == oldColor {
		fill(image, m, n, nextSr, sc, oldColor, newColor)
	}
	nextSc := sc - 1
	if 0 <= nextSc && nextSc < n && image[sr][nextSc] == oldColor {
		fill(image, m, n, sr, nextSc, oldColor, newColor)
	}
	nextSr = sr + 1
	if 0 <= nextSr && nextSr < m && image[nextSr][sc] == oldColor {
		fill(image, m, n, nextSr, sc, oldColor, newColor)
	}
	nextSc = sc + 1
	if 0 <= nextSc && nextSc < n && image[sr][nextSc] == oldColor {
		fill(image, m, n, sr, nextSc, oldColor, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	n := len(image[0])
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	fill(image, m, n, sr, sc, oldColor, newColor)
	return image
}
