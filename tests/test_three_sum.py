from solutions.three_sum import Solution

def test_example_1():
    solver = Solution()
    nums = [-1, 0, 1, 2, -1, -4]
    actual = solver.threeSum(nums)
    actual_sorted = sorted([sorted(t) for t in actual])
    expected = [[-1, -1, 2], [-1, 0, 1]]
    expected_sorted = sorted([sorted(t) for t in expected])
    assert actual_sorted == expected_sorted

def test_example_2():
    solver = Solution()
    assert solver.threeSum([0, 1, 1]) == []

def test_example_3():
    solver = Solution()
    assert solver.threeSum([0, 0, 0]) == [[0, 0, 0]]

def test_empty():
    solver = Solution()
    assert solver.threeSum([]) == []

def test_no_solution():
    solver = Solution()
    assert solver.threeSum([1, 2, 3]) == []
