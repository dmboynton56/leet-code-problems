from __future__ import annotations


class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None):
        self.val = val
        self.next = next


def solution(head: ListNode | None) -> bool:
    if head is None or head.next is None:
        return False

    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True

    return False


if __name__ == "__main__":
    n3 = ListNode(0)
    n2 = ListNode(2, n3)
    n1 = ListNode(2, n2)
    head = ListNode(3, n1)
    n3.next = n1
    print(solution(head))  # True

    print(solution(ListNode(1)))  # False
