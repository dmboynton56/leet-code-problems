class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        charSet = set()
        l = 0
        ans = 0
        for i in range(len(s)):
            while s[i] in charSet: #invalid, move left pointer to the right
                charSet.remove(s[l])
                l += 1
            charSet.add(s[i])
            ans = max(ans, i - l + 1)
        return ans
