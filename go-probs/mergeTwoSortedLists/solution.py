from __future__ import annotations


class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None):
        self.val = val
        self.next = next


def solution(list1: ListNode | None, list2: ListNode | None) -> ListNode | None:
    """
    Dummy-head merge. Python uses None for end-of-list; Go uses nil *ListNode.
    """
    dummy = ListNode()
    curr = dummy

    while list1 and list2:
        if list1.val <= list2.val:
            curr.next = list1
            list1 = list1.next
        else:
            curr.next = list2
            list2 = list2.next
        curr = curr.next

    curr.next = list1 if list1 else list2
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

    print(to_list(solution(build([1, 2, 4]), build([1, 3, 4]))))
