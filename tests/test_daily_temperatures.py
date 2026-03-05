from solutions.daily_temperatures import Solution

def test_example_1():
    solver = Solution()
    temperatures = [73, 74, 75, 71, 69, 72, 76, 73]
    assert solver.dailyTemperatures(temperatures) == [1, 1, 4, 2, 1, 1, 0, 0]

def test_example_2():
    solver = Solution()
    temperatures = [30, 40, 50, 60]
    assert solver.dailyTemperatures(temperatures) == [1, 1, 1, 0]

def test_example_3():
    solver = Solution()
    temperatures = [30, 60, 90]
    assert solver.dailyTemperatures(temperatures) == [1, 1, 0]

def test_decreasing():
    solver = Solution()
    temperatures = [90, 80, 70]
    assert solver.dailyTemperatures(temperatures) == [0, 0, 0]

def test_empty():
    solver = Solution()
    assert solver.dailyTemperatures([]) == []
