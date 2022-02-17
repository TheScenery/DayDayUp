package floodfill

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2
	r := floodFill(image, sr, sc, newColor)
	fmt.Println(r)

	image = [][]int{{0, 0, 0}, {0, 0, 0}}
	sr = 0
	sc = 0
	newColor = 2
	r = floodFill(image, sr, sc, newColor)
	fmt.Println(r)
}
