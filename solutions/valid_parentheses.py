from typing import List

class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        closeToOpen = {")": "(", "]": "[", "}": "{"}

        for ch in s:
            if ch in closeToOpen: #closing
                if stack and stack[-1] == closeToOpen[ch]:
                    stack.pop()
                else:
                    return False
            else: #opening
                stack.append(ch)
        return True if not stack else False