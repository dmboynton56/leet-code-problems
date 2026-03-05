from solutions.container_with_most_water import Solution

def test_example_1():
    solver = Solution()
    assert solver.maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49

def test_example_2():
    solver = Solution()
    assert solver.maxArea([1, 1]) == 1

def test_decreasing():
    solver = Solution()
    assert solver.maxArea([4, 3, 2, 1]) == 4 # (4, 1) -> min(4, 1) * 3 = 3; (4, 2) -> min(4, 2) * 2 = 4

def test_increasing():
    solver = Solution()
    assert solver.maxArea([1, 2, 3, 4]) == 4
