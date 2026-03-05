from solutions.group_anagrams import Solution

def test_example_1():
    solver = Solution()
    strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
    actual = solver.groupAnagrams(strs)
    # Sorting each inner list and the outer list to compare
    actual_sorted = sorted([sorted(group) for group in actual])
    expected = [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
    expected_sorted = sorted([sorted(group) for group in expected])
    assert actual_sorted == expected_sorted

def test_example_2():
    solver = Solution()
    strs = [""]
    assert solver.groupAnagrams(strs) == [[""]]

def test_example_3():
    solver = Solution()
    strs = ["a"]
    assert solver.groupAnagrams(strs) == [["a"]]

def test_no_anagrams():
    solver = Solution()
    strs = ["abc", "def", "ghi"]
    actual = solver.groupAnagrams(strs)
    actual_sorted = sorted([sorted(group) for group in actual])
    expected = [["abc"], ["def"], ["ghi"]]
    expected_sorted = sorted([sorted(group) for group in expected])
    assert actual_sorted == expected_sorted
