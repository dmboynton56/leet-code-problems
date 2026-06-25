def solution(nums: list[int], target: int) -> list[int]:
    """
    One-pass hash map. Python dicts grow dynamically; no make() needed.
    Go uses map[int]int and the `val, ok := m[key]` idiom instead of `in`.
    """
    seen: dict[int, int] = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i
    return []


if __name__ == "__main__":
    print(solution([2, 7, 11, 15], 9))  # [0, 1]
    print(solution([3, 2, 4], 6))       # [1, 2]
    print(solution([3, 3], 6))          # [0, 1]
