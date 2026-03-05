class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        charSet = set()
        l = 0
        ans = 0
        for i in range(len(s)):
            print("i, l, charSet, s[i]", i, l, charSet, s[i])
            while s[i] in charSet: #invalid, move left pointer to the right
                charSet.remove(s[l])
                l += 1
            charSet.add(s[i])
            ans = max(ans, i - l + 1)
        return ans
