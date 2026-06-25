def solution(nums: list[int]) -> list[list[int]]:
    """
    Sort + two pointers. Python sorted() returns new list; Go sort.Slice sorts in place.
    """
    nums.sort()
    result: list[list[int]] = []

    for i in range(len(nums) - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue

        left, right = i + 1, len(nums) - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]
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
