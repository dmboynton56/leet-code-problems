from solutions.meeting_rooms import Solution

def test_example_1():
    solver = Solution()
    intervals = [[0, 30], [5, 10], [15, 20]]
    assert solver.canAttendMeetings(intervals) is False

def test_example_2():
    solver = Solution()
    intervals = [[7, 10], [2, 4]]
    assert solver.canAttendMeetings(intervals) is True

def test_overlap_exact_edge():
    # Typically [1, 2] and [2, 3] do not overlap in this problem context
    solver = Solution()
    intervals = [[1, 2], [2, 3]]
    assert solver.canAttendMeetings(intervals) is True

def test_empty():
    solver = Solution()
    assert solver.canAttendMeetings([]) is True

def test_single_meeting():
    solver = Solution()
    assert solver.canAttendMeetings([[1, 2]]) is True
