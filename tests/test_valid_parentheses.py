# tests/test_valid_parentheses.py

# 1. Import the Solution class you want to test
from solutions.valid_parentheses import Solution

# 2. Define a test function for the first LeetCode example
def test_example_1():
    # Arrange: Set up the inputs and expected output
    solver = Solution()
    s = "()"
    expected_output = True

    # Act: Call the method you're testing
    actual_output = solver.isValid(s)

    # Assert: Check if the actual result matches the expected one
    assert actual_output == expected_output


# 3. Define another test function for the second example
def test_example_2():
    # Arrange
    solver = Solution()
    s = "()[]{}"
    expected_output = True

    # Act
    actual_output = solver.isValid(s)

    # Assert
    assert actual_output == expected_output


# You can add more test cases for edge scenarios
def test_example_3():
    """LeetCode example: invalid pairing."""
    solver = Solution()
    s = "(]"
    assert solver.isValid(s) is False


def test_example_4():
    """LeetCode example: wrong nesting order."""
    solver = Solution()
    s = "([)]"
    assert solver.isValid(s) is False


def test_example_5():
    """LeetCode example: valid nested brackets."""
    solver = Solution()
    s = "{[]}"
    assert solver.isValid(s) is True


def test_empty_string():
    """Empty string is considered valid."""
    solver = Solution()
    s = ""
    assert solver.isValid(s) is True


def test_single_open_bracket():
    """Single bracket can't be closed."""
    solver = Solution()
    s = "("
    assert solver.isValid(s) is False


def test_single_close_bracket():
    """Closing bracket without open is invalid."""
    solver = Solution()
    s = "]"
    assert solver.isValid(s) is False


def test_odd_length_string():
    """Odd length can't be fully paired."""
    solver = Solution()
    s = "(()"
    assert solver.isValid(s) is False