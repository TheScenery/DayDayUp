# Title:3. Longest Substring Without Repeating Characters
# Description:
# Given a string, find the length of the longest substring without repeating characters.
# Example:
# Given "abcabcbb", the answer is "abc", which the length is 3.
# Given "bbbbb", the answer is "b", with the length of 1.
# Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.



# Language: Python

class Solution:
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        indexMaping = {}
        startIndex = 0
        longLength = 0
        for i in range(0, len(s)):
            char = s[i]
            if char in indexMaping:
                tempLength = i - startIndex
                if tempLength > longLength:
                    longLength = tempLength
                if indexMaping[char] + 1 > startIndex:
                    startIndex = indexMaping[char] + 1
                indexMaping[char] = i
            else:
                indexMaping[char] = i
        tempLength = len(s) - startIndex
        if tempLength > longLength:
            longLength = tempLength
        return longLength


solution = Solution()
print(solution.lengthOfLongestSubstring("abcabcbb"))  