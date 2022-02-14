package singleelementinasortedarray

func singleNonDuplicate(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	mid := length / 2
	if mid%2 != 0 {
		mid -= 1
	}
	if nums[mid] != nums[mid+1] {
		return singleNonDuplicate(nums[0 : mid+1])
	}
	return singleNonDuplicate(nums[mid+2:])
}
