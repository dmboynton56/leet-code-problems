# Given a string s, return the index of the first non-repeating character. If none exists, return -1.
# Example: "leetcode" -> 0, "loveleetcode" -> 2, "aabb" -> -1.

def non_repeating_character(s: str) -> int:
    freqMap = {}
    if not s: return -1
    for i in range(len(s)):
        if s[i] not in freqMap:
            freqMap[s[i]] = 1
        else:
            freqMap[s[i]]+= 1
    # for ch in freqMap:
    #     if freqMap[ch] == 1:
    #         repeatingChar = ch
    #         break
    for i in range(len(s)):
        if freqMap[s[i]] == 1:
            return i
    return -1

s = "imsofuckingstupidlmfao"
print(non_repeating_character(s))