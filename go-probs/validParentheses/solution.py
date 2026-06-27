# solution checks whether brackets are properly opened and closed.
# Go: iterate runes (Unicode code points), not bytes — important for non-ASCII.
# Python: iterating a str also yields single-character strings; here ASCII-only is fine.
def solution(s: str) -> bool:
    stack: list[str] = []  # Go slice used as stack; Python uses list with append/pop.
    pairs = {")": "(", "]": "[", "}": "{"}  # Python shortcut: one map close→open; Go uses explicit if chains per char.

    for char in s:
        # range over string yields runes in Go; `char` is rune (int32), not byte.
        if char in "([{":
            stack.append(char)  # append returns new slice in Go; reassign is idiomatic. Python: stack.append()
        else:
            # Combines Go's len(stack)==0 guard and top-mismatch checks in one expression.
            if not stack or stack[-1] != pairs[char]:
                return False  # Go bool is explicit; no truthiness on empty collections like Python.
            stack.pop()  # Go: stack = stack[:len(stack)-1]; Python: stack.pop()

    return len(stack) == 0


if __name__ == "__main__":
    print(solution("()"))       # True
    print(solution("()[]{}"))   # True
    print(solution("(]"))       # False
    print(solution("([)]"))    # False
    print(solution("{[]}"))     # True
