package twosummodule

func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums {
			if v1+v2 == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	for i, v1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v1+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	hash := map[int]int{}
	for i, v1 := range nums {
		p, ok := hash[target-v1]
		if ok {
			return []int{p, i}
		}
		hash[v1] = i
	}
	return nil
}
