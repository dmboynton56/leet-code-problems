from solutions.permutation_in_string import Solution

def test_example_1():
    solver = Solution()
    assert solver.checkInclusion("ab", "eidbaooo") is True

def test_example_2():
    solver = Solution()
    assert solver.checkInclusion("ab", "eidboaoo") is False

def test_exact_match():
    solver = Solution()
    assert solver.checkInclusion("adc", "dcda") is True

def test_s1_longer_than_s2():
    solver = Solution()
    assert solver.checkInclusion("hello", "ooolleoooleh") is False
