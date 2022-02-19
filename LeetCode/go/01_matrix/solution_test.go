package matrix01

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	r := updateMatrix(mat)
	fmt.Println(r)

	mat = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	r = updateMatrix(mat)
	fmt.Println(r)
}

func TestFunc1(t *testing.T) {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	r := updateMatrix1(mat)
	fmt.Println(r)

	mat = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	r = updateMatrix1(mat)
	fmt.Println(r)
}
