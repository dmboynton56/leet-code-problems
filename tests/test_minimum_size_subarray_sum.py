from solutions.minimum_size_subarray_sum import Solution

def test_example_1():
    solver = Solution()
    assert solver.minSubArrayLen(7, [2, 3, 1, 2, 4, 3]) == 2

def test_example_2():
    solver = Solution()
    assert solver.minSubArrayLen(4, [1, 4, 4]) == 1

def test_example_3():
    solver = Solution()
    assert solver.minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1]) == 0

def test_single_element_exact():
    solver = Solution()
    assert solver.minSubArrayLen(5, [5]) == 1

def test_single_element_less():
    solver = Solution()
    assert solver.minSubArrayLen(5, [4]) == 0
