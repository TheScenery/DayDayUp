from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        k = len(nums) - 1
        i = 0
        while i <= k:
            if nums[i] == val:
                while k >= 0 and nums[k] == val:
                    k -= 1
                if i < k:
                    temp = nums[i]
                    nums[i] = nums[k]
                    nums[k] = temp
                    k -= 1
            i += 1

        return k + 1


s = Solution()
nums1 = [0, 4, 4, 0, 4, 4, 4, 0, 2]
print(s.removeElement(nums1, 4))
print(nums1)
