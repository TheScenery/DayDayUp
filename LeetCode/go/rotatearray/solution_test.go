package rotatearray

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func TestRotate1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate1(nums, k)
	fmt.Println(nums)
}
