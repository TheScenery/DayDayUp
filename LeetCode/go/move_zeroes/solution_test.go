package movezeros

import (
	"fmt"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func TestMoveZeros1(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes1(nums)
	fmt.Println(nums)
}

func TestMoveZeros2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes2(nums)
	fmt.Println(nums)
}
