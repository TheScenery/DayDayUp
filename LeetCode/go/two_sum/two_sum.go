package twosummodule

func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums {
			if v1+v2 == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
