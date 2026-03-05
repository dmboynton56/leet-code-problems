from solutions.insert_interval import Solution

def test_example_1():
    solver = Solution()
    intervals = [[1, 3], [6, 9]]
    newInterval = [2, 5]
    assert solver.insert(intervals, newInterval) == [[1, 5], [6, 9]]

def test_example_2():
    solver = Solution()
    intervals = [[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]]
    newInterval = [4, 8]
    assert solver.insert(intervals, newInterval) == [[1, 2], [3, 10], [12, 16]]

def test_no_overlap_before():
    solver = Solution()
    intervals = [[3, 5]]
    newInterval = [1, 2]
    assert solver.insert(intervals, newInterval) == [[1, 2], [3, 5]]

def test_no_overlap_after():
    solver = Solution()
    intervals = [[1, 2]]
    newInterval = [3, 5]
    assert solver.insert(intervals, newInterval) == [[1, 2], [3, 5]]

def test_empty():
    solver = Solution()
    assert solver.insert([], [1, 2]) == [[1, 2]]
