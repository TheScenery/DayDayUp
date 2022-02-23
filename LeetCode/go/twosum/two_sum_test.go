package twosummodule

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	if result[0] != 0 || result[1] != 1 {
		t.Fail()
	}

	result = twoSum([]int{3, 2, 4}, 6)
	if result[0] != 1 || result[1] != 2 {
		t.Fail()
	}
}

func TestTwoSun3(t *testing.T) {
	result := twoSum3([]int{2, 7, 11, 15}, 9)
	if result[0] != 0 || result[1] != 1 {
		t.Fail()
	}

	result = twoSum3([]int{3, 2, 4}, 6)
	if result[0] != 1 || result[1] != 2 {
		t.Fail()
	}
}
