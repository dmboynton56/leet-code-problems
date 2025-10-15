# tests/test_summary_ranges.py

# Import the Solution class from where your solution is saved
from solutions.summary_ranges import Solution

def test_example_1():
    """Tests the first example from the problem description."""
    # Arrange
    solver = Solution()
    nums = [0, 1, 2, 4, 5, 7]
    expected = ["0->2", "4->5", "7"]
    
    # Act
    actual = solver.summaryRanges(nums)
    
    # Assert
    assert actual == expected

def test_example_2():
    """Tests the second example from the problem description."""
    # Arrange
    solver = Solution()
    nums = [0, 2, 3, 4, 6, 8, 9]
    expected = ["0", "2->4", "6", "8->9"]
    
    # Act
    actual = solver.summaryRanges(nums)
    
    # Assert
    assert actual == expected

def test_empty_list():
    """Tests an empty input list, which should return an empty list."""
    # Arrange
    solver = Solution()
    nums = []
    expected = []
    
    # Act
    actual = solver.summaryRanges(nums)
    
    # Assert
    assert actual == expected

def test_single_element():
    """Tests a list with only one number."""
    # Arrange
    solver = Solution()
    nums = [5]
    expected = ["5"]
    
    # Act
    actual = solver.summaryRanges(nums)
    
    # Assert
    assert actual == expected

def test_all_consecutive_numbers():
    """Tests a list where all numbers form a single range."""
    # Arrange
    solver = Solution()
    nums = [0, 1, 2, 3, 4, 5]
    expected = ["0->5"]
    
    # Act
    actual = solver.summaryRanges(nums)
    
    # Assert
    assert actual == expected

def test_no_consecutive_numbers():
    """Tests a list where no numbers are consecutive."""
    # Arrange
    solver = Solution()
    nums = [1, 3, 5, 7, 9]
    expected = ["1", "3", "5", "7", "9"]
    
    # Act
    actual = solver.summaryRanges(nums)
    
    # Assert
    assert actual == expected

def test_with_negative_numbers():
    """Tests a list that includes negative numbers."""
    # Arrange
    solver = Solution()
    nums = [-5, -4, -2, 0, 1, 3]
    expected = ["-5->-4", "-2", "0->1", "3"]
    
    # Act
    actual = solver.summaryRanges(nums)
    
    # Assert
    assert actual == expected