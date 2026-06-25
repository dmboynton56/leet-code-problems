def solution(n: int) -> int:
    """
    Fibonacci DP. Python allows elegant tuple swap; Go uses explicit prev1/prev2 ints.
    """
    if n <= 2:
        return n

    prev2, prev1 = 1, 2
    for _ in range(3, n + 1):
        prev2, prev1 = prev1, prev2 + prev1

    return prev1


if __name__ == "__main__":
    print(solution(2))  # 2
    print(solution(3))  # 3
    print(solution(5))  # 8
