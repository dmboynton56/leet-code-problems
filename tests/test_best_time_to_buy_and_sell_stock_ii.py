from solutions.best_time_to_buy_and_sell_stock_ii import Solution

def test_example_1():
    solver = Solution()
    assert solver.maxProfit([7, 1, 5, 3, 6, 4]) == 7

def test_example_2():
    solver = Solution()
    assert solver.maxProfit([1, 2, 3, 4, 5]) == 4

def test_example_3():
    solver = Solution()
    assert solver.maxProfit([7, 6, 4, 3, 1]) == 0

def test_single_day():
    solver = Solution()
    assert solver.maxProfit([5]) == 0
