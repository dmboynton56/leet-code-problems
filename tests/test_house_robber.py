from solutions.house_robber import Solution

def test_example_1():
    solver = Solution()
    assert solver.rob([1, 2, 3, 1]) == 4

def test_example_2():
    solver = Solution()
    assert solver.rob([2, 7, 9, 3, 1]) == 12

def test_empty():
    solver = Solution()
    assert solver.rob([]) == 0

def test_single_house():
    solver = Solution()
    assert solver.rob([5]) == 5
