package nextgreaterelementi

import "container/list"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	for i, n := range nums1 {
		result[i] = -1
		startFind := false
		for _, n2 := range nums2 {
			if n2 == n {
				startFind = true
			}
			if startFind && n2 > n {
				result[i] = n2
				break
			}
		}
	}
	return result
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	hashTable := make(map[int]int)
	l := list.New()
	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]
		back := l.Back()
		if back == nil {
			l.PushBack(n)
			hashTable[n] = -1
		} else {
			backValue := back.Value.(int)
			if n < backValue {
				l.PushBack(n)
				hashTable[n] = backValue
			} else {
				for back != nil && n > back.Value.(int) {
					l.Remove(back)
					back = l.Back()
				}
				if back == nil {
					hashTable[n] = -1
					l.PushBack(n)
				} else {
					hashTable[n] = back.Value.(int)
					l.PushBack(n)
				}
			}
		}
	}
	result := make([]int, len(nums1))
	for i, n := range nums1 {
		result[i] = hashTable[n]
	}
	return result
}
