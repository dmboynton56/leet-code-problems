# Linked List Cycle (Easy) — LeetCode #141

**Difficulty:** Easy  
**Pattern:** Floyd's cycle detection (fast/slow pointers)

## Data Structure

- **Singly linked list** with possible cycle back to some node.
- **Two pointers** — slow moves 1 step, fast moves 2 steps per iteration.

## Algorithm

1. If head is nil or head.Next is nil, no cycle.
2. Initialize `slow, fast = head, head`.
3. While `fast != nil && fast.Next != nil`:
   - Advance slow by 1, fast by 2.
   - If `slow == fast`, cycle exists → return true.
4. Fast reached end → return false.

**Time:** O(n)  
**Space:** O(1)

## Edge Cases

| Case | Notes |
|------|-------|
| Empty list | false |
| Single node, no cycle | false |
| Cycle to head | Detected when pointers meet |
| Cycle to middle node | Same algorithm |
| Self-loop on one node | fast catches slow |

## Go-Specific Notes

- **ListNode struct** — typically `type ListNode struct { Val int; Next *ListNode }`.
- **Pointer comparison** — `slow == fast` compares addresses; valid for cycle detection.
- Alternative: `map[*ListNode]bool` visited set — O(n) space, simpler but less optimal.

## Other Notes

- Linked List Cycle II (#142) asks for the cycle entry node (phase 2 of Floyd's).
- Same fast/slow pattern used in finding middle of list.
