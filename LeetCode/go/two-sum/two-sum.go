package twosum

// description: https://leetcode.cn/problems/two-sum
func twoSum(nums []int, target int) []int {
	tempMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if targetIndex, ok := tempMap[nums[i]]; ok {
			return []int{i, targetIndex}
		}
		tempMap[target-nums[i]] = i
	}
	return nil
}
