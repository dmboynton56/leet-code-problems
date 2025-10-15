# tests/test_number_of_islands.py

# 1. Import the Solution class you want to test
from solutions.number_of_islands import Solution

# 2. Define a test function for the first LeetCode example
def test_example_1():
    # Arrange: Set up the inputs and expected output
    solver = Solution()
    grid = [
      ["1","1","1","1","0"],
      ["1","1","0","1","0"],
      ["1","1","0","0","0"],
      ["0","0","0","0","0"]
    ]
    expected_output = 1
    
    # Act: Call the method you're testing
    actual_output = solver.numIslands(grid)
    
    # Assert: Check if the actual result matches the expected one
    assert actual_output == expected_output

# 3. Define another test function for the second example
def test_example_2():
    # Arrange
    solver = Solution()
    grid = [
      ["1","1","0","0","0"],
      ["1","1","0","0","0"],
      ["0","0","1","0","0"],
      ["0","0","0","1","1"]
    ]
    expected_output = 3
    
    # Act
    actual_output = solver.numIslands(grid)

    # Assert
    assert actual_output == expected_output

# You can add more test cases for edge scenarios
def test_empty_grid():
    """Tests an empty grid, which should result in 0 islands."""
    solver = Solution()
    grid = []
    assert solver.numIslands(grid) == 0

def test_no_islands():
    """Tests a grid with only water."""
    solver = Solution()
    grid = [
        ["0", "0", "0"],
        ["0", "0", "0"]
    ]
    assert solver.numIslands(grid) == 0