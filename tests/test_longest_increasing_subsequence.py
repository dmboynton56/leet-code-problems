from solutions.longest_increasing_subsequence import Solution

def test_example_1():
    solver = Solution()
    assert solver.lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18]) == 4

def test_example_2():
    solver = Solution()
    assert solver.lengthOfLIS([0, 1, 0, 3, 2, 3]) == 4

def test_example_3():
    solver = Solution()
    assert solver.lengthOfLIS([7, 7, 7, 7, 7, 7, 7]) == 1

def test_empty():
    solver = Solution()
    assert solver.lengthOfLIS([]) == 0
