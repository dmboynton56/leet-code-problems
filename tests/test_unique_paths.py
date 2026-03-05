from solutions.unique_paths import Solution

def test_example_1():
    solver = Solution()
    assert solver.uniquePaths(3, 7) == 28

def test_example_2():
    solver = Solution()
    assert solver.uniquePaths(3, 2) == 3

def test_minimal_grid():
    solver = Solution()
    assert solver.uniquePaths(1, 1) == 1
