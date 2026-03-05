from solutions.minimum_path_sum import Solution

def test_example_1():
    solver = Solution()
    grid = [[1, 3, 1], [1, 5, 1], [4, 2, 1]]
    assert solver.minPathSum(grid) == 7

def test_example_2():
    solver = Solution()
    grid = [[1, 2, 3], [4, 5, 6]]
    assert solver.minPathSum(grid) == 12

def test_single_cell():
    solver = Solution()
    assert solver.minPathSum([[5]]) == 5
