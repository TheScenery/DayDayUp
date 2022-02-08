package medianoftwosortedarrays

import (
	"sort"
)

func medianOfSortedArray(nums []int) float64 {
	length := len(nums)
	if length%2 == 0 {
		return float64((nums[length/2] + nums[length/2-1])) / 2
	}
	return float64(nums[(length-1)/2])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	return medianOfSortedArray(nums)
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	lengthNums1 := len(nums1)
	lengthNums2 := len(nums2)
	if lengthNums1 == 0 && lengthNums2 == 0 {
		return 0
	}
	if lengthNums1 == 0 {
		return medianOfSortedArray(nums2)
	}
	if lengthNums2 == 0 {
		return medianOfSortedArray(nums1)
	}
	nums := make([]int, lengthNums1+lengthNums2)
	i, j, k := 0, 0, 0

	for {
		if i >= lengthNums1 && j >= lengthNums2 {
			break
		}
		if i >= lengthNums1 {
			nums[k] = nums2[j]
			j++
		} else if j >= lengthNums2 {
			nums[k] = nums1[i]
			i++
		} else {
			num1 := nums1[i]
			num2 := nums2[j]
			if num2 < num1 {
				nums[k] = num2
				j++
			} else {
				nums[k] = num1
				i++
			}
		}
		k++
	}

	return medianOfSortedArray(nums)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func getKth(nums1, nums2 []int, k int) int {
	lengthNums1 := len(nums1)
	lengthNums2 := len(nums2)
	if lengthNums1 == 0 {
		return nums2[k-1]
	}
	if lengthNums2 == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	half := k / 2
	i := min(lengthNums1-1, half-1)
	j := min(lengthNums2-1, half-1)
	x := nums1[i]
	y := nums2[j]
	if x <= y {
		nums1 = nums1[i+1:]
		k = k - i - 1
	} else {
		nums2 = nums2[j+1:]
		k = k - j - 1
	}
	return getKth(nums1, nums2, k)
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	half := length / 2
	if length%2 == 0 {
		return float64(getKth(nums1, nums2, half)+getKth(nums1, nums2, half+1)) / 2
	}
	return float64(getKth(nums1, nums2, half+1))
}
