# solution returns all unique triplets that sum to zero.
# Go: sort.Slice with less func; Python: nums.sort() or sorted().
def solution(nums: list[int]) -> list[list[int]]:
    nums.sort()  # Go sort.Slice sorts in place; Python sorted() returns new list, sort() is in-place.
    result: list[list[int]] = []
    n = len(nums)

    for i in range(n - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue  # skip duplicate anchors; Python same guard

        left, right = i + 1, n - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]  # Go names this `sum`
            if total == 0:
                result.append([nums[i], nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and nums[left] == nums[left - 1]:
                    left += 1
                while left < right and nums[right] == nums[right + 1]:
                    right -= 1
            elif total < 0:
                left += 1
            else:
                right -= 1

    return result


if __name__ == "__main__":
    print(solution([-1, 0, 1, 2, -1, -4]))
    # [[-1, -1, 2], [-1, 0, 1]]
