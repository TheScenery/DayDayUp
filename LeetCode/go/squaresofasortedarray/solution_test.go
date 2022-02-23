package squaresofasortedarray

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	r := sortedSquares(nums)
	fmt.Println(r)
}

func TestSortedSquares1(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	r := sortedSquares1(nums)
	fmt.Println(r)

	nums = []int{-7, -3, 2, 3, 11}
	r = sortedSquares1(nums)
	fmt.Println(r)
}
