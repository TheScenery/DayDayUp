package partitionequalsubsetsum

import "testing"

func TestFunc(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	r := canPartition(nums)
	if r != true {
		t.Fatal(r)
	}

	nums = []int{1, 2, 3, 5}
	r = canPartition(nums)
	if r != false {
		t.Fatal(r)
	}
}

func TestFunc1(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	r := canPartition1(nums)
	if r != true {
		t.Fatal(r)
	}

	nums = []int{1, 2, 3, 5}
	r = canPartition1(nums)
	if r != false {
		t.Fatal(r)
	}
}
