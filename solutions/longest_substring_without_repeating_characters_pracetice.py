class Solution:    
    def longestSubstring(self, s: str) -> int:
        charSet = set()
        res = 0
        l = 0
        for i in range(len(s)):
            while s[i] in charSet: #invalid
                charSet.remove(s[l])
                l += 1
            charSet.add(s[i])
            res = max(res, i - l + 1)
        return res
