package medianoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	if result != float64(2) {
		t.Fatal()
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	result = findMedianSortedArrays(nums1, nums2)
	if result != 2.5 {
		t.Fatal()
	}
}

func TestFindMedianSortedArrays1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays1(nums1, nums2)
	if result != float64(2) {
		t.Fatal()
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	result = findMedianSortedArrays1(nums1, nums2)
	if result != 2.5 {
		t.Fatal()
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays2(nums1, nums2)
	if result != float64(2) {
		t.Fatal()
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	result = findMedianSortedArrays2(nums1, nums2)
	if result != 2.5 {
		t.Fatal()
	}

	nums1 = []int{1}
	nums2 = []int{2, 3, 4, 5, 6}
	result = findMedianSortedArrays2(nums1, nums2)
	if result != 3.5 {
		t.Fatal()
	}
}

func TestGetKth(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	result := getKth(nums1, nums2, 2)
	if result != 2 {
		t.Fatal()
	}
	result = getKth(nums1, nums2, 3)
	if result != 3 {
		t.Fatal()
	}

	nums1 = []int{1}
	nums2 = []int{2, 3, 4, 5, 6}
	result = getKth(nums1, nums2, 3)
	if result != 3 {
		t.Fatal()
	}
	result = getKth(nums1, nums2, 4)
	if result != 4 {
		t.Fatal()
	}
}
