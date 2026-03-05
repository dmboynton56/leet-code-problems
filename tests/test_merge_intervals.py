from solutions.merge_intervals import Solution

def test_example_1():
    solver = Solution()
    intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
    assert solver.merge(intervals) == [[1, 6], [8, 10], [15, 18]]

def test_example_2():
    solver = Solution()
    intervals = [[1, 4], [4, 5]]
    assert solver.merge(intervals) == [[1, 5]]

def test_no_overlap():
    solver = Solution()
    intervals = [[1, 2], [3, 4]]
    assert solver.merge(intervals) == [[1, 2], [3, 4]]

def test_empty():
    solver = Solution()
    assert solver.merge([]) == []

def test_single_interval():
    solver = Solution()
    assert solver.merge([[1, 2]]) == [[1, 2]]
