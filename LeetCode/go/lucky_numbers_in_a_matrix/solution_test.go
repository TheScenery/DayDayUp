package luckynumbersinamatrix

import (
	"fmt"
	"testing"
)

func TestLuckyNumbers(t *testing.T) {
	matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	r := luckyNumbers(matrix)
	fmt.Println(r)

	matrix = [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	r = luckyNumbers(matrix)
	fmt.Println(r)

	matrix = [][]int{{7, 8}, {1, 2}}
	r = luckyNumbers(matrix)
	fmt.Println(r)
}
