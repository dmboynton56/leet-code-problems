# tests/test_longest_substring_without_repeating_characters.py

# 1. Import the Solution class you want to test
from solutions.longest_substring_without_repeating_characters import Solution

# 2. Define a test function for the first LeetCode example
def test_example_1():
    # Arrange: Set up the inputs and expected output
    solver = Solution()
    s = "abcabcbb"
    expected_output = 3

    # Act: Call the method you're testing
    actual_output = solver.lengthOfLongestSubstring(s)

    # Assert: Check if the actual result matches the expected one
    assert actual_output == expected_output


# 3. Define another test function for the second example
def test_example_2():
    # Arrange
    solver = Solution()
    s = "bbbbb"
    expected_output = 1

    # Act
    actual_output = solver.lengthOfLongestSubstring(s)

    # Assert
    assert actual_output == expected_output


def test_example_3():
    """LeetCode example: repeats inside, longest length is 3 ("wke")."""
    solver = Solution()
    s = "pwwkew"
    assert solver.lengthOfLongestSubstring(s) == 3


# You can add more test cases for edge scenarios
def test_empty_string():
    """Empty string should return 0."""
    solver = Solution()
    s = ""
    assert solver.lengthOfLongestSubstring(s) == 0


def test_single_character():
    """Single character should return 1."""
    solver = Solution()
    s = "a"
    assert solver.lengthOfLongestSubstring(s) == 1


def test_all_unique_characters():
    """All unique characters => answer is the full length."""
    solver = Solution()
    s = "abcdef"
    assert solver.lengthOfLongestSubstring(s) == 6


def test_string_with_spaces():
    """Spaces count as characters."""
    solver = Solution()
    s = " "
    assert solver.lengthOfLongestSubstring(s) == 1


def test_tricky_case_dvdf():
    """Common tricky input where answer is 3 ('vdf')."""
    solver = Solution()
    s = "dvdf"
    assert solver.lengthOfLongestSubstring(s) == 3


def test_tricky_case_anviaj():
    """Common tricky input where answer is 5 ('nviaj')."""
    solver = Solution()
    s = "anviaj"
    assert solver.lengthOfLongestSubstring(s) == 5