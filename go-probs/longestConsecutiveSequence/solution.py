def solution(nums: list[int]) -> int:
    seen = set(nums)
    max_len = 0

    for num in seen:
        if num - 1 in seen:
            continue
        length = 1
        curr = num + 1
        while curr in seen:
            length += 1
            curr += 1
        max_len = max(max_len, length)

    return max_len


if __name__ == "__main__":
    print(solution([100, 4, 200, 1, 3, 2]))              # 4
    print(solution([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]))    # 9
