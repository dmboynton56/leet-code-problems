from solutions.valid_palindrome import Solution

def test_example_1():
    solver = Solution()
    assert solver.isPalindrome("A man, a plan, a canal: Panama") is True

def test_example_2():
    solver = Solution()
    assert solver.isPalindrome("race a car") is False

def test_example_3():
    solver = Solution()
    assert solver.isPalindrome(" ") is True

def test_single_char():
    solver = Solution()
    assert solver.isPalindrome("a") is True

def test_numbers():
    solver = Solution()
    assert solver.isPalindrome("0P") is False
    assert solver.isPalindrome("121") is True
