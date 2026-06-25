def solution(nums: list[int]) -> list[int]:
    """
    Prefix + suffix in O(1) extra space (output array excluded).
    Python can use result list; Go uses make([]int, n).
    """
    n = len(nums)
    result = [1] * n

    prefix = 1
    for i in range(n):
        result[i] = prefix
        prefix *= nums[i]

    suffix = 1
    for i in range(n - 1, -1, -1):
        result[i] *= suffix
        suffix *= nums[i]

    return result


if __name__ == "__main__":
    print(solution([1, 2, 3, 4]))
    print(solution([-1, 1, 0, -3, 3]))
