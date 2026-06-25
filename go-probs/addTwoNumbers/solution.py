from __future__ import annotations


class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None):
        self.val = val
        self.next = next


def solution(l1: ListNode | None, l2: ListNode | None) -> ListNode | None:
    """
    Digit-by-digit add with carry. Python: sum // 10 and sum % 10, same as Go.
    """
    dummy = ListNode()
    curr = dummy
    carry = 0

    while l1 or l2 or carry:
        total = carry
        if l1:
            total += l1.val
            l1 = l1.next
        if l2:
            total += l2.val
            l2 = l2.next
        curr.next = ListNode(total % 10)
        curr = curr.next
        carry = total // 10

    return dummy.next


if __name__ == "__main__":
    def build(vals: list[int]) -> ListNode | None:
        dummy = ListNode()
        c = dummy
        for v in vals:
            c.next = ListNode(v)
            c = c.next
        return dummy.next

    def to_list(head: ListNode | None) -> list[int]:
        out = []
        while head:
            out.append(head.val)
            head = head.next
        return out

    print(to_list(solution(build([2, 4, 3]), build([5, 6, 4]))))
