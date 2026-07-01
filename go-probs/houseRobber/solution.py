def solution(nums: list[int]) -> int:
    prev, curr = 0, 0
    for num in nums:
        prev, curr = curr, max(curr, prev + num)
    return curr


if __name__ == "__main__":
    print(solution([1, 2, 3, 1]))      # 4
    print(solution([2, 7, 9, 3, 1]))  # 12
