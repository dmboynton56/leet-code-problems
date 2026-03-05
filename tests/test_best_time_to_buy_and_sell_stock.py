from solutions.best_time_to_buy_and_sell_stock import Solution

def test_example_1():
    solver = Solution()
    assert solver.maxProfit([7, 1, 5, 3, 6, 4]) == 5

def test_example_2():
    solver = Solution()
    assert solver.maxProfit([7, 6, 4, 3, 1]) == 0

def test_empty():
    solver = Solution()
    assert solver.maxProfit([]) == 0

def test_single_price():
    solver = Solution()
    assert solver.maxProfit([10]) == 0

def test_monotonically_increasing():
    solver = Solution()
    assert solver.maxProfit([1, 2, 3, 4, 5]) == 4
