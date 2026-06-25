def solution(nums: list[int]) -> int:
    """
    Kadane's algorithm. Python one-liner variant: current = max(x, current + x).
    """
    max_sum = current = nums[0]

    for num in nums[1:]:
        current = max(num, current + num)
        max_sum = max(max_sum, current)

    return max_sum


if __name__ == "__main__":
    print(solution([-2, 1, -3, 4, -1, 2, 1, -5, 4]))  # 6
    print(solution([1]))                                  # 1
    print(solution([5, 4, -1, 7, 8]))                   # 23
