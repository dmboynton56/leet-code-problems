from solutions.decode_ways import Solution

def test_example_1():
    solver = Solution()
    assert solver.numDecodings("12") == 2

def test_example_2():
    solver = Solution()
    assert solver.numDecodings("226") == 3

def test_example_3():
    solver = Solution()
    assert solver.numDecodings("06") == 0

def test_single_zero():
    solver = Solution()
    assert solver.numDecodings("0") == 0
