package sortanarray

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	r := sortArray(nums)
	fmt.Println(r)

	nums = []int{5, 1, 1, 2, 0, 0}
	r = sortArray(nums)
	fmt.Println(r)
}

func TestSelection(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	r := sortArray1(nums)
	fmt.Println(r)

	nums = []int{5, 1, 1, 2, 0, 0}
	r = sortArray1(nums)
	fmt.Println(r)
}

func TestQuickSort(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	r := sortArray2(nums)
	fmt.Println(r)

	nums = []int{5, 1, 1, 2, 0, 0}
	r = sortArray2(nums)
	fmt.Println(r)
}
