# solution returns length of longest substring without repeating characters.
# Go: map[rune]int for last index; Python dict works similarly.
def solution(s: str) -> int:
    last_seen: dict[str, int] = {}  # Go: make(map[rune]int)
    max_len = 0
    start = 0  # left bound of sliding window (inclusive index in rune slice sense)

    for i, char in enumerate(s):
        # range index i is byte offset for ASCII but use map on rune; for ASCII s, i aligns with rune index.
        if char in last_seen and last_seen[char] >= start:
            # Go: if prev, ok := lastSeen[char]; ok && prev >= start
            start = last_seen[char] + 1  # shrink window from left; Python same logic
        last_seen[char] = i
        max_len = max(max_len, i - start + 1)  # Go: if window := i-start+1; window > maxLen

    return max_len


if __name__ == "__main__":
    print(solution("abcabcbb"))  # 3
    print(solution("bbbbb"))     # 1
    print(solution("pwwkew"))    # 3
