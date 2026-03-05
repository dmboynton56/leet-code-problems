from solutions.product_of_array_except_self import Solution

def test_example_1():
    solver = Solution()
    assert solver.productExceptSelf([1, 2, 3, 4]) == [24, 12, 8, 6]

def test_example_2():
    solver = Solution()
    assert solver.productExceptSelf([-1, 1, 0, -3, 3]) == [0, 0, 9, 0, 0]

def test_zeros():
    solver = Solution()
    assert solver.productExceptSelf([0, 0]) == [0, 0]
    assert solver.productExceptSelf([1, 0]) == [0, 1]
