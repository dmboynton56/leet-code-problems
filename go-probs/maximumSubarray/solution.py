# solution returns the largest sum of any contiguous subarray (Kadane's algorithm).
# Go: track running sum and global max; Python's max() is built-in.
def solution(nums: list[int]) -> int:
    max_sum = current = nums[0]  # assume non-empty per problem

    for num in nums[1:]:
        # extend current subarray or start fresh at nums[i]
        # Python one-liner via max(); Go uses explicit if currentSum+nums[i] > nums[i]
        current = max(num, current + num)
        max_sum = max(max_sum, current)

    return max_sum


if __name__ == "__main__":
    print(solution([-2, 1, -3, 4, -1, 2, 1, -5, 4]))  # 6
    print(solution([1]))                                  # 1
    print(solution([5, 4, -1, 7, 8]))                   # 23
