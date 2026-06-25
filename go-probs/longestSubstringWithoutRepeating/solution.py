def solution(s: str) -> int:
    """
    Sliding window with last-seen index map.
    Python for-loop index matches char position; Go range on string uses byte offsets
    (same as char index for pure ASCII strings like LeetCode examples).
    """
    last_seen: dict[str, int] = {}
    max_len = 0
    start = 0

    for i, char in enumerate(s):
        if char in last_seen and last_seen[char] >= start:
            start = last_seen[char] + 1
        last_seen[char] = i
        max_len = max(max_len, i - start + 1)

    return max_len


if __name__ == "__main__":
    print(solution("abcabcbb"))  # 3
    print(solution("bbbbb"))     # 1
    print(solution("pwwkew"))    # 3
