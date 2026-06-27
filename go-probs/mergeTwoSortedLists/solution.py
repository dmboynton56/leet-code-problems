from __future__ import annotations


# ListNode matches LeetCode's linked list definition.
# Go: struct with pointer next; Python often uses a class or nested ListNode in LeetCode stub.
class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None):
        self.val = val
        self.next = next


# solution merges two sorted linked lists into one sorted list.
# Go: pointer manipulation; dummy head avoids special-casing the first node.
def solution(list1: ListNode | None, list2: ListNode | None) -> ListNode | None:
    dummy = ListNode()  # heap-allocated via & in Go; Python objects are always references.
    curr = dummy

    while list1 and list2:  # Go: list1 != nil && list2 != nil
        if list1.val <= list2.val:
            curr.next = list1
            list1 = list1.next  # advance pointer; Python: list1 = list1.next
        else:
            curr.next = list2
            list2 = list2.next
        curr = curr.next

    # Attach remainder — at most one is non-nil.
    # Python ternary: one line vs Go's explicit if/else on list1 != nil.
    curr.next = list1 if list1 else list2

    return dummy.next  # skip sentinel node


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

    print(to_list(solution(build([1, 2, 4]), build([1, 3, 4]))))  # [1, 2, 3, 4, 4]
