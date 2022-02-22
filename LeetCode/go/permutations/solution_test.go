package permutations

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	nums := []int{1, 2, 3}
	r := permute(nums)
	fmt.Println(r)
}
