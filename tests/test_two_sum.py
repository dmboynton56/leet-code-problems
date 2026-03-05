from solutions.two_sum import Solution

def test_example_1():
    solver = Solution()
    nums = [2, 7, 11, 15]
    target = 9
    expected = [0, 1]
    # The order of the indices doesn't matter, but standard is sorted
    assert sorted(solver.twoSum(nums, target)) == sorted(expected)

def test_example_2():
    solver = Solution()
    nums = [3, 2, 4]
    target = 6
    expected = [1, 2]
    assert sorted(solver.twoSum(nums, target)) == sorted(expected)

def test_example_3():
    solver = Solution()
    nums = [3, 3]
    target = 6
    expected = [0, 1]
    assert sorted(solver.twoSum(nums, target)) == sorted(expected)

def test_no_solution():
    # According to problem constraints, there's always exactly one solution,
    # but it's good to handle edge cases in testing if needed.
    pass

def test_negative_numbers():
    solver = Solution()
    nums = [-3, 4, 3, 90]
    target = 0
    expected = [0, 2]
    assert sorted(solver.twoSum(nums, target)) == sorted(expected)
