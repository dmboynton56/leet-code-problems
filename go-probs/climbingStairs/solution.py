# solution counts distinct ways to climb n stairs (1 or 2 steps at a time).
# Go: iterative DP with two ints; Python often uses a loop or @lru_cache recursion.
def solution(n: int) -> int:
    if n <= 2:
        return n  # base cases: 1 way for n=1, 2 ways for n=2

    prev2, prev1 = 1, 2  # fibonacci-style: ways(n-2), ways(n-1)
    for _ in range(3, n + 1):
        # Go: curr := prev1+prev2; prev2=prev1; prev1=curr
        # Python tuple swap: rotate window in one line.
        prev2, prev1 = prev1, prev2 + prev1

    return prev1


if __name__ == "__main__":
    print(solution(2))  # 2
    print(solution(3))  # 3
    print(solution(5))  # 8
