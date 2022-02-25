package wherewilltheballfall

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	grid := [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
	r := findBall(grid)
	fmt.Println(r)

	grid = [][]int{{-1}}
	r = findBall(grid)
	fmt.Println(r)

	grid = [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}
	r = findBall(grid)
	fmt.Println(r)
}
