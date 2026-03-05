from solutions.top_k_frequent_elements import Solution

def test_example_1():
    solver = Solution()
    nums = [1, 1, 1, 2, 2, 3]
    k = 2
    actual = solver.topKFrequent(nums, k)
    assert sorted(actual) == [1, 2]

def test_example_2():
    solver = Solution()
    nums = [1]
    k = 1
    assert solver.topKFrequent(nums, k) == [1]

def test_different_counts():
    solver = Solution()
    nums = [4, 1, -1, 2, -1, 2, 3]
    k = 2
    actual = solver.topKFrequent(nums, k)
    assert sorted(actual) == [-1, 2]
