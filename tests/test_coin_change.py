from solutions.coin_change import Solution

def test_example_1():
    solver = Solution()
    assert solver.coinChange([1, 2, 5], 11) == 3

def test_example_2():
    solver = Solution()
    assert solver.coinChange([2], 3) == -1

def test_example_3():
    solver = Solution()
    assert solver.coinChange([1], 0) == 0

def test_no_solution():
    solver = Solution()
    assert solver.coinChange([2, 5], 3) == -1
