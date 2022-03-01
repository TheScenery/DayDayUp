package nextgreaterelementi

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	r := nextGreaterElement(nums1, nums2)
	fmt.Println(r)

	nums1 = []int{2, 4}
	nums2 = []int{1, 2, 3, 4}
	r = nextGreaterElement(nums1, nums2)
	fmt.Println(r)

	nums1 = []int{4, 1, 2}
	nums2 = []int{1, 2, 3, 4}
	r = nextGreaterElement(nums1, nums2)
	fmt.Println(r)
}

func TestFunc1(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	r := nextGreaterElement1(nums1, nums2)
	fmt.Println(r)

	nums1 = []int{2, 4}
	nums2 = []int{1, 2, 3, 4}
	r = nextGreaterElement1(nums1, nums2)
	fmt.Println(r)

	nums1 = []int{4, 1, 2}
	nums2 = []int{1, 2, 3, 4}
	r = nextGreaterElement1(nums1, nums2)
	fmt.Println(r)

	nums1 = []int{1, 3, 5, 2, 4}
	nums2 = []int{6, 5, 4, 3, 2, 1, 7}
	r = nextGreaterElement1(nums1, nums2)
	fmt.Println(r)
}
