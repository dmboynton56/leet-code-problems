# solution finds two indices whose values sum to target.
# Go: explicit types, map[int]int, and the comma-ok idiom for map lookups.
# Python equivalent: dict with `if complement in seen` (no ok tuple).
def solution(nums: list[int], target: int) -> list[int]:
    seen: dict[int, int] = {}  # Go maps must be initialized with make() or a literal; nil maps panic on write.

    for i, num in enumerate(nums):
        # range returns (index, value). Python: for i, num in enumerate(nums)
        complement = target - num
        if complement in seen:
            # Go has no ternary; return a slice literal. Python: return [seen[complement], i]
            return [seen[complement], i]
        seen[num] = i

    return []  # Go returns zero-value slice; Python might return [] or raise.


if __name__ == "__main__":
    print(solution([2, 7, 11, 15], 9))  # [0, 1]
    print(solution([3, 2, 4], 6))       # [1, 2]
    print(solution([3, 3], 6))          # [0, 1]
