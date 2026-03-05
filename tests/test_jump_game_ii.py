from solutions.jump_game_ii import Solution

def test_example_1():
    solver = Solution()
    assert solver.jump([2, 3, 1, 1, 4]) == 2

def test_example_2():
    solver = Solution()
    assert solver.jump([2, 3, 0, 1, 4]) == 2

def test_single_element():
    solver = Solution()
    assert solver.jump([0]) == 0
