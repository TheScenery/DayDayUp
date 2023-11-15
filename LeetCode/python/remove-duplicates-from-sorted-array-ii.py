from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return len(nums)
        i = 0
        j = 1
        currentVal = nums[i]
        count = 1
        while j < len(nums):
            if currentVal == nums[j]:
                if count < 2:
                    i += 1
                    count += 1
                    nums[i] = nums[j]
            else:
                currentVal = nums[j]
                i += 1
                nums[i] = nums[j]
                count = 1
            j += 1

        return i + 1


s = Solution()
nums1 = [1, 1, 1, 2, 2, 3]
print(s.removeDuplicates(nums1))
print(nums1)
