def solution(height: list[int]) -> int:
    left, right = 0, len(height) - 1
    max_area = 0

    while left < right:
        h = min(height[left], height[right])
        width = right - left
        max_area = max(max_area, h * width)
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1

    return max_area


if __name__ == "__main__":
    print(solution([1, 8, 6, 2, 5, 4, 8, 3, 7]))  # 49
    print(solution([1, 1]))                       # 1
