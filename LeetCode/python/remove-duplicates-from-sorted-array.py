from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        i = j = 0
        while j < len(nums):
            if i == j:
                j += 1
            elif nums[i] == nums[j]:
                j += 1
            else:
                i += 1
                nums[i] = nums[j]

        return i + 1


s = Solution()
nums1 = [1, 1, 2]
print(s.removeDuplicates(nums1))
print(nums1)
