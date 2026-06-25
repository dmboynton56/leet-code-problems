from collections import Counter

def non_repeating_character(s: str) -> int:
    if not s:
        return -1
    
    # Counter creates the frequency map in one line
    count = Counter(s)
    
    # Get index and character at the same time (i, char in enumerate(str))
    # Second pass to find the first character with a frequency of 1
    for i, char in enumerate(s):
        print(char, "char", count[char], "count")
        if count[char] == 1:
            return i
            
    return -1

# Examples
print(non_repeating_character("leetcode"))      # 0
print(non_repeating_character("loveleetcode"))  # 2
print(non_repeating_character("aabb"))          # -1