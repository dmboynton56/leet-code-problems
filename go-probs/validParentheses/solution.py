def solution(s: str) -> bool:
    """
    Stack-based bracket matching.
    Python list as stack: append/pop. Go uses []rune slice with append/reslice.
    """
    stack: list[str] = []
    pairs = {')': '(', ']': '[', '}': '{'}

    for char in s:
        if char in '([{':
            stack.append(char)
        else:
            if not stack or stack[-1] != pairs[char]:
                return False
            stack.pop()

    return len(stack) == 0


if __name__ == "__main__":
    print(solution("()"))       # True
    print(solution("()[]{}"))   # True
    print(solution("(]"))       # False
    print(solution("([)]"))    # False
    print(solution("{[]}"))     # True
