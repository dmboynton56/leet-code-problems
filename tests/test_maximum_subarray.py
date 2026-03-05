from solutions.maximum_subarray import Solution

def test_example_1():
    solver = Solution()
    assert solver.maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]) == 6

def test_example_2():
    solver = Solution()
    assert solver.maxSubArray([1]) == 1

def test_example_3():
    solver = Solution()
    assert solver.maxSubArray([5, 4, -1, 7, 8]) == 23

def test_all_negative():
    solver = Solution()
    assert solver.maxSubArray([-5, -1, -8]) == -1
