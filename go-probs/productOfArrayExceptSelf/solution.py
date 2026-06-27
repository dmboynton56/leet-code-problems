# solution returns product of all elements except self without division.
# Go: two passes with output slice; prefix/suffix products.
def solution(nums: list[int]) -> list[int]:
    n = len(nums)
    result = [1] * n  # make allocates zeroed slice in Go; Python [1]*n then overwrite

    # Forward pass: result[i] = product of nums[0..i-1]
    # Go folds prefix into result[i-1]; Python keeps a separate `prefix` variable (same math).
    prefix = 1
    for i in range(n):
        result[i] = prefix
        prefix *= nums[i]

    # Backward pass: multiply by suffix product on the fly
    suffix = 1
    for i in range(n - 1, -1, -1):
        result[i] *= suffix
        suffix *= nums[i]

    return result


if __name__ == "__main__":
    print(solution([1, 2, 3, 4]))  # [24, 12, 8, 6]
    print(solution([-1, 1, 0, -3, 3]))
